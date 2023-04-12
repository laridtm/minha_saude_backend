db = new Mongo().getDB("minha_saude");

db.records.insertMany([
    {
        "_id": "6435fd7aa0683320460d5dbc",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "Clinica dos olhos",
        "Professional": "Dr. João",
        "Observation": "Consulta oftalmologista",
        "Type": "appointment"
    },
    {
        "_id": "6435fd8ba0683320460d5dbd",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "Unidade básica de saúde",
        "Professional": "Dr. José",
        "Observation": "Vacina contra gripe",
        "Type": "vaccine"
    },
    {
        "_id": "6435fd8ba0683320460d5dbe",
        "userid": "00897314921",
        "Date": ISODate("2023-04-23T19:00:00Z"),
        "Hospital": "São Lucas",
        "Professional": "Dra. Rafaela",
        "Observation": "Exame de sangue",
        "Type": "exam"
    }
]);