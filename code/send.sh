curl \
    -X POST \
    -u key_123: \
    -d '{
        "template": "tem_123",
        "recipient": {
            "address": "greg@sendwithus.com"
        },
        "template_data": {
            "first_name": "Greg"
        }
    }' \
    https://api.sendwithus.com/api/v1/send
