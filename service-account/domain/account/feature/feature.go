package feature

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"service-account/config"
	"service-account/domain/account/constant"
	"service-account/domain/account/helper"
	"service-account/domain/account/model"
	repository "service-account/domain/account/repository"
	shared_constant "service-account/domain/shared/constant"
	Error "service-account/domain/shared/error"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
)

type AccountFeature interface {
	CreateRekeningFeature(ctx context.Context, request model.CreateAccountRequest) (response model.NasabahJSON, err error)
	SavingFeature(ctx context.Context, request model.SavingAccountRequest) (response model.SaldoJSON, err error)
	WitdrawalFeature(ctx context.Context, request model.WitdrawalAccountRequest) (response model.SaldoJSON, err error)
	BalanceFeature(ctx context.Context, noRekening string) (response model.SaldoJSON, err error)
	HistoryFeature(ctx context.Context, noRekening string) (response []model.HistoryJSON, err error)
	TransferFeature(ctx context.Context, request model.TransferRequest) (response model.SaldoJSON, err error)
	GenerateTokenFeature(ctx context.Context, request model.GenerateToken) (response model.Token, err error)
	ShortRunningScript(ctx context.Context)
}

type accountFeature struct {
	config            config.EnvironmentConfig
	accountRepository repository.AccountRepository
	redis             *redis.Client
}

func NewAccountFeature(config config.EnvironmentConfig, accountRepository repository.AccountRepository, redis *redis.Client) AccountFeature {
	return &accountFeature{
		config:            config,
		accountRepository: accountRepository,
		redis:             redis,
	}
}

func (t *accountFeature) CreateRekeningFeature(ctx context.Context, request model.CreateAccountRequest) (response model.NasabahJSON, err error) {

	// check NIK
	checkNIK, _ := t.accountRepository.GetByNIK(ctx, request.NIK)

	if checkNIK != nil && checkNIK.ID != 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNIKAlreadyExist, errors.New(fmt.Sprintf(constant.ErrNIKAlreadyExistWithNIK, request.NIK)))
		return
	}

	// check NoHP
	checkNoHP, _ := t.accountRepository.GetByNoHandphone(ctx, request.NoHandphone)

	if checkNoHP != nil && checkNoHP.ID != 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoHandphoneAlreadyExist, errors.New(fmt.Sprintf(constant.ErrNoHandphoneAlreadyExistWithNoHP, request.NoHandphone)))
		return
	}

	// Hash Password
	hashPassword, errHash := helper.HashPassword(request.PIN)
	if errHash != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrPin, errors.New(fmt.Sprintf(constant.ErrPinWithNoPIN, request.PIN)))
		return
	}

	// Fill Data
	data := model.Account{
		Nama:        request.Nama,
		NIK:         request.NIK,
		PIN:         hashPassword,
		NoHandphone: request.NoHandphone,
		NoRekening:  helper.GenerateNoRekening(request),
		Saldo:       int64(constant.DEFAULT_BALANCE),
	}

	err = t.accountRepository.CreateRekeningRepository(ctx, &data)
	if err != nil {
		return
	}

	response = model.NasabahJSON{
		NoRekening: data.NoRekening,
	}
	return
}

func (t *accountFeature) SavingFeature(ctx context.Context, request model.SavingAccountRequest) (response model.SaldoJSON, err error) {
	// Get Data
	getNorek, _ := t.accountRepository.GetByNoRekening(ctx, request.NoRekening)

	if getNorek.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, request.NoRekening)))
		return
	}

	// Fill Data
	data := model.Account{
		ID:         getNorek.ID,
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo + request.Nominal,
	}

	err = t.accountRepository.UpdateSaldoRepository(ctx, &data)
	if err != nil {
		return
	}

	// Fill Data Mutasi
	mutasi := model.History{
		NoRekening:    data.NoRekening,
		KodeTransaksi: constant.CREDIT,
		Nominal:       request.Nominal,
		Saldo:         data.Saldo,
	}

	err = t.accountRepository.CreateHistoryRepository(ctx, &mutasi)
	if err != nil {
		return
	}

	response = model.SaldoJSON{
		NoRekening: data.NoRekening,
		Saldo:      data.Saldo,
	}

	t.SendJournal(ctx, model.Journal{
		TanggalTramsaksi: time.Now(),
		NoRekeningKredit: data.NoRekening,
		NoRekeningDebit:  0,
		NominalKredit:    request.Nominal,
		NominalDebit:     0,
	})

	return
}

func (t *accountFeature) WitdrawalFeature(ctx context.Context, request model.WitdrawalAccountRequest) (response model.SaldoJSON, err error) {
	// Get Data
	getNorek, _ := t.accountRepository.GetByNoRekening(ctx, request.NoRekening)

	if getNorek.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, request.NoRekening)))
		return
	}

	// Validate Saldo
	if request.Nominal > getNorek.Saldo {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrSaldo, errors.New(fmt.Sprint(constant.ErrSaldo)))
		return
	}

	// Fill Data
	data := model.Account{
		ID:         getNorek.ID,
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo - request.Nominal,
	}

	err = t.accountRepository.UpdateSaldoRepository(ctx, &data)
	if err != nil {
		return
	}

	// Fill Data Mutasi
	mutasi := model.History{
		NoRekening:    data.NoRekening,
		KodeTransaksi: constant.DEBIT,
		Nominal:       request.Nominal,
		Saldo:         data.Saldo,
	}

	err = t.accountRepository.CreateHistoryRepository(ctx, &mutasi)
	if err != nil {
		return
	}

	response = model.SaldoJSON{
		NoRekening: data.NoRekening,
		Saldo:      data.Saldo,
	}

	go t.SendJournal(ctx, model.Journal{
		TanggalTramsaksi: time.Now(),
		NoRekeningKredit: 0,
		NoRekeningDebit:  data.NoRekening,
		NominalKredit:    0,
		NominalDebit:     request.Nominal,
	})
	return
}

func (t *accountFeature) BalanceFeature(ctx context.Context, noRekening string) (response model.SaldoJSON, err error) {
	getNorek, _ := t.accountRepository.GetByNoRekening(ctx, helper.StrToInt64(noRekening))

	if getNorek.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, noRekening)))
		return
	}

	response = model.SaldoJSON{
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo,
	}
	return
}

func (t *accountFeature) HistoryFeature(ctx context.Context, noRekening string) (response []model.HistoryJSON, err error) {
	getNorek, _ := t.accountRepository.GetByNoRekening(ctx, helper.StrToInt64(noRekening))
	if getNorek.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, noRekening)))
		return
	}

	getData, err := t.accountRepository.GetHistoryRepository(ctx, getNorek.ID)
	if err != nil {
		return
	}

	for _, v := range getData {
		response = append(response, model.HistoryJSON{
			KodeTransaksi: v.KodeTransaksi,
			Nominal:       v.Nominal,
			Saldo:         v.Saldo,
			Waktu:         v.CreatedAt,
		})
	}
	return
}

func (t *accountFeature) TransferFeature(ctx context.Context, request model.TransferRequest) (response model.SaldoJSON, err error) {
	// Validate Norek
	if request.NoRekeningAsal == request.NoRekeningTujuan {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrSendSameAccount, errors.New(fmt.Sprint(constant.ErrSendSameAccount)))
		return
	}

	// Get Norek Asal
	getNorekAsal, _ := t.accountRepository.GetByNoRekening(ctx, request.NoRekeningAsal)

	if getNorekAsal.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrYourNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrYourNoRekeningNotFoundWithNoRekening, request.NoRekeningAsal)))
		return
	}

	// Get Norek Tujuan
	getNorekTujuan, _ := t.accountRepository.GetByNoRekening(ctx, request.NoRekeningTujuan)

	if getNorekTujuan.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningDestinationNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningDestinationNotFoundWithNoRekening, request.NoRekeningTujuan)))
		return
	}

	// Validate Saldo
	if request.Nominal > getNorekAsal.Saldo {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrSaldo, errors.New(fmt.Sprint(constant.ErrSaldo)))
		return
	}

	// Fill Data Asal
	dataAsal := model.Account{
		ID:         getNorekAsal.ID,
		NoRekening: getNorekAsal.NoRekening,
		Saldo:      getNorekAsal.Saldo - request.Nominal,
	}

	err = t.accountRepository.UpdateSaldoRepository(ctx, &dataAsal)
	if err != nil {
		return
	}

	// Fill Data Mutasi Asal
	mutasiAsal := model.History{
		NoRekening:    dataAsal.NoRekening,
		KodeTransaksi: constant.TRANSFER,
		Nominal:       request.Nominal,
		Saldo:         dataAsal.Saldo,
	}

	err = t.accountRepository.CreateHistoryRepository(ctx, &mutasiAsal)
	if err != nil {
		return
	}

	// Fill Data Tujuan
	dataTujuan := model.Account{
		ID:         getNorekTujuan.ID,
		NoRekening: getNorekTujuan.NoRekening,
		Saldo:      getNorekTujuan.Saldo + request.Nominal,
	}

	err = t.accountRepository.UpdateSaldoRepository(ctx, &dataTujuan)
	if err != nil {
		return
	}

	// Fill Data Mutasi Tujuan
	mutasiTujuan := model.History{
		NoRekening:    dataTujuan.NoRekening,
		KodeTransaksi: constant.TRANSFER,
		Nominal:       request.Nominal,
		Saldo:         dataTujuan.Saldo,
	}

	err = t.accountRepository.CreateHistoryRepository(ctx, &mutasiTujuan)
	if err != nil {
		return
	}

	response = model.SaldoJSON{
		NoRekening: dataAsal.NoRekening,
		Saldo:      dataAsal.Saldo,
	}

	t.SendJournal(ctx, model.Journal{
		TanggalTramsaksi: time.Now(),
		NoRekeningKredit: dataTujuan.NoRekening,
		NoRekeningDebit:  dataAsal.NoRekening,
		NominalKredit:    request.Nominal,
		NominalDebit:     request.Nominal,
	})

	return
}

func (t *accountFeature) SendJournal(ctx context.Context, request model.Journal) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error serializing data:", err)
		return
	}
	err = t.redis.Set(ctx, "insert-data-key", string(jsonData), 1*time.Second).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}
}

func (t *accountFeature) GenerateTokenFeature(ctx context.Context, request model.GenerateToken) (response model.Token, err error) {
	// Get Data
	getNorek, _ := t.accountRepository.GetByNoRekening(ctx, request.NoRekening)

	if getNorek.ID == 0 {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, request.NoRekening)))
		return
	}

	if !helper.CheckPasswordHash(request.PIN, getNorek.PIN) {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrPinNotFound, errors.New(fmt.Sprint(constant.ErrPinNotFound)))
		return
	}

	secretKey := "secret-key"

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"username": getNorek.NoRekening,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	response.Token = tokenString

	return
}

func (t *accountFeature) ShortRunningScript(ctx context.Context) {
	for {
		var (
			jumlahNominalTarik    int64
			jumlahNominalSetor    int64
			jumlahNominalTransfer int64
		)

		datas, _ := t.accountRepository.GetAllHistoryRepository(ctx)
		if len(datas) != 0 {
			for _, data := range datas {
				if data.KodeTransaksi == "C" {
					jumlahNominalSetor += data.Nominal
				} else if data.KodeTransaksi == "D" {
					jumlahNominalTarik += data.Nominal
				} else if data.KodeTransaksi == "T" {
					jumlahNominalTransfer += data.Nominal
				}
			}
		}
		fmt.Println()
		fmt.Println("-------------- START SHORT SCRIPT -------------")
		fmt.Printf("Total Transaksi :%v ", len(datas))
		fmt.Println()
		fmt.Printf("Jumlah Nominal Tarik :%v ", jumlahNominalTarik)
		fmt.Println()
		fmt.Printf("Jumlah Nominal Setor :%v ", jumlahNominalSetor)
		fmt.Println()
		fmt.Printf("Jumlah Nominal Transfer :%v ", jumlahNominalTransfer)
		fmt.Println()
		fmt.Println("-------------- END SHORT SCRIPT ----------------")
		time.Sleep(1 * time.Minute)
	}

}
