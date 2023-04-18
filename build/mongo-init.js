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

db.records.insertMany([
    {
        "_id": "6435fd7aa0683320460d5dbc",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "Clinica dos olhos",
        "Professional": "Dr. João",
        "Name": "Clínico geral",
        "Observation": "Consulta oftalmologista",
        "Type": "appointment"
    },
    {
        "_id": "6435fd8ba0683320460d5dbd",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "Unidade básica de saúde",
        "Professional": "Dr. José",
        "Name": "Hepatite A",
        "Observation": "Vacina contra gripe",
        "Type": "vaccine"
    },
    {
        "_id": "6435fd8ba0683320460d5dbe",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "São Lucas",
        "Professional": "Dra. Rafaela",
        "Name": "Raio X",
        "Observation": "Exame de sangue",
        "Type": "exam"
    }
]);