#!/bin/bash

mongo --version
mongo <<EOF
var config = {
    "_id": "dbrs",
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "localhost:27017",
            "priority": 1
        }
    ]
};
rs.initiate(config, { force: true });
rs.status();
use Cadastro;
db.createCollection("Filial")