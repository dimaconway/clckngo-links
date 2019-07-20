let collectionName = 'links';

db.createCollection(collectionName, {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['createdDatetime', 'url', 'code', 'hits', 'lastHitDatetime'],
            properties: {
                createdDatetime: {
                    bsonType: 'date',
                },
                url: {
                    bsonType: 'string',
                },
                code: {
                    bsonType: 'string',
                },
                hits: {
                    bsonType: 'int',
                },
                lastHitDatetime: {
                    bsonType: 'date',
                }
            }
        }
    },
    validationLevel: 'strict',
    validationAction: 'error'
});

db[collectionName].createIndex(
    {
        code: 1
    },
    {
        unique: true
    }
);
