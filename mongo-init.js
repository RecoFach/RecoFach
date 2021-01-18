print('Start #################################################################');

db = db.getSiblingDB('core-db');
db.createUser(
  {
    user: 'admin',
    pwd: 'admin',
    roles: [{ role: 'readWrite', db: 'core-db' }],
  },
);


print('END #################################################################');
