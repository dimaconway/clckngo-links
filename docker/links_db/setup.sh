#mongo=( mongo --host 127.0.0.1 --port 27017 --quiet )
mongo=( mongo )
mongo+=(
    '--username='"$MONGO_INITDB_ROOT_USERNAME"
    '--password='"$MONGO_INITDB_ROOT_PASSWORD"
    '--authenticationDatabase='"$MONGO_ROOT_AUTH_DB"
)

"${mongo[@]}" "$MONGO_INITDB_DATABASE" --eval "db.createUser({user:'$MONGO_USER',pwd:'$MONGO_PASSWORD',roles:[{role:'readWrite', db:'$MONGO_INITDB_DATABASE'}]})"

CREATE_FILES=( /scripts/*-create.js )
for f in "${CREATE_FILES[@]}"; do "${mongo[@]}" "$MONGO_INITDB_DATABASE" $f; done
