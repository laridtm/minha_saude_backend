db = new Mongo().getDB("minha_saude");

db.profile.insertMany([
    {
        "_id": "643ad8a8259c06d923b5c2ad",
        "fullName": "Larissa Diniz Martins",
        "birthDate": "20/08/1997",
        "cpf": "00897314921",
        "phoneNumber": "(48) 99652-5859",
        "address": "Rua Roberto Sell, 264",
        "maritalStatus": "Casada",
        "bloodType": "O+",
        "emergencyPhone": "(48) 99836-9511",
        "allergies": "Penicilina, Novalgina"
    }
])

db.reminders.insertMany([
    {
        "_id": "6442b054cdb79df5d192d547",
        "name": "Remédio Pressão",
        "time": "12:00",
        "userId": "00897314921",
        "type": "everyDay"
    },
    {
        "_id": "6442b11831249fbe90ba9de5",
        "name": "Consulta Cardio",
        "time": "14:00",
        "userId": "00897314921",
        "type": "once"
    },
    {
        "_id": "6442b11931249fbe90ba9de7",
        "name": "Remédio Gastrite",
        "time": "21:00",
        "userId": "00897314921",
        "type": "everyDay"
    },
    {
        "_id": "6442b11a31249fbe90ba9de8",
        "name": "Vitaminas",
        "time": "08:30",
        "userId": "00897314921",
        "type": "weekends"
    },
])

db.records.insertMany([
    {
        "_id": "6435fd7aa0683320460d5dbc",
        "userId": "00897314921",
        "date": ISODate("2023-04-23T19:00:00Z"),
        "hospital": "Clinica dos olhos",
        "professional": "Dr. João",
        "name": "Clínico geral",
        "observation": "Consulta oftalmologista",
        "type": "appointment"
    },
    {
        "_id": "6435fd8ba0683320460d5dbd",
        "userId": "00897314921",
        "date": ISODate("2023-04-23T19:00:00Z"),
        "hospital": "Unidade básica de saúde",
        "professional": "Dr. José",
        "name": "Hepatite A",
        "observation": "Vacina contra gripe",
        "type": "vaccine"
    },
    {
        "_id": "6435fd8ba0683320460d5dbe",
        "userId": "00897314921",
        "date": ISODate("2023-04-23T19:00:00Z"),
        "hospital": "São Lucas",
        "professional": "Dra. Rafaela",
        "name": "Raio X",
        "observation": "Exame de sangue",
        "type": "exam"
    }
]);