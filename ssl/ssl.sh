# to check the ssl maped to the domain

openssl x509 -in /etc/vpl/ssl/fullchain.pem -noout -text | grep -E "Subject:|DNS:"


#  To get dokcer images id  


docker images --format="{{.ID}}"

# To delete the docker images at once use this

docker rmi $(docker images --format="{{.ID}}")