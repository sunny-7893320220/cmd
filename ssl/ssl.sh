# to check the ssl maped to the domain

openssl x509 -in /etc/vpl/ssl/fullchain.pem -noout -text | grep -E "Subject:|DNS:"