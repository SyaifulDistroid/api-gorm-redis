package constant

const (
	ErrNIKAlreadyExist        = "nik already exist"
	ErrNIKAlreadyExistWithNIK = "nik %s already exist"

	ErrNoHandphoneAlreadyExist         = "no handphone already exist"
	ErrNoHandphoneAlreadyExistWithNoHP = "no handphone %s already exist"

	ErrNoRekeningNotFound               = "no rekening not found"
	ErrNoRekeningNotFoundWithNoRekening = "no rekening %v not found"

	ErrPinNotFound = "no rekening and pin not match"

	ErrYourNoRekeningNotFound               = "your no rekening not found"
	ErrYourNoRekeningNotFoundWithNoRekening = "your no rekening %v not found"

	ErrSendSameAccount = "cannot send to the same account"

	ErrNoRekeningDestinationNotFound               = "destination no rekening not found"
	ErrNoRekeningDestinationNotFoundWithNoRekening = "destination no rekening %v not found"

	ErrSaldo = "the balance is insufficient"

	ErrPin          = "no pin format wrong"
	ErrPinWithNoPIN = "no pin %s format wrong"
)
