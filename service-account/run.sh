
API_URL="http://127.0.0.1:8787/v1/account/tabung"

BEARER_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY4MzA5NzYsInVzZXJuYW1lIjoyNTQyNzU5MH0.Vk_YJ5Cyc8wA4jHdVv-wpgc7_dNnFQdwOO5fzpsrCqc"

JSON_PAYLOAD='{"nomor_rekening": 25427590, "nominal": 1000}'

curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $BEARER_TOKEN" \
    -d "$JSON_PAYLOAD" \
    "$API_URL"
