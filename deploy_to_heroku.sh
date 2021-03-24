go mod vendor
echo "Type the commit message: "
read -r commitMessage
git add .
git commit -am "${commitMessage}"
heroku container:push -a whats-api-cli web
heroku container:release -a whats-api-cli web