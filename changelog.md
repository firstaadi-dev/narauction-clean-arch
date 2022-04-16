# Barang Struct

## Changes:

- idEvent -> eventID
- asal_daerah -> asalDaerah
- ada createdAt,updatedAt, deletedAt tapi gakepake

## Prev:

```
{
  "id": "barang-E5LzZ27KAMO6t9ku",
  "lot": "Lot-10",
  "namaBarang": "Motif Semen Rante",
  "foto": ["https://i.ibb.co/QcZPknY/b-10-1.jpg"],
  "tahunPembuatan": 1990,
  "descId": "Motif ini meniru Batik Kraton. Motif  live auction on the Discord Narauction platform.",
  "priceRange": [2100000, 99999999],
  "size": [236, 103],
  "tipe": "Katun Lokal",
  "idEvent": "event-tlHqWFQwSH2a2fKK",
  "asal_daerah": "Solo",
  "dyeType": "\"\"",
  "hargaAwal": 1000000,
  "urlThumbnail": [""],
  "isAvailable": false
}

```

## New:

```
{
  "id": 2,
  "createdAt": "0001-01-01T00:00:00Z",
  "updatedAt": "0001-01-01T00:00:00Z",
  "deletedAt": null,
  "lot": "Lot-06",
  "namaBarang": "Motif Kotak Gurdha Grinsing Klithik (Batik Kraton)",
  "foto": [
    "https://i.ibb.co/rQQKTdt/14-1.jpg",
    "https://i.ibb.co/jWjdg97/14-2.jpg",
    "https://i.ibb.co/m67L4d9/14-3.jpg",
    "https://i.ibb.co/486NKgX/14-4.jpg",
    "https://i.ibb.co/bKf7HzF/14-5.jpg"
  ],
  "tahunPembuatan": 1960,
  "descId": "Motif ini merupakan motif pengembangan yang ada didalam tembok Keraton.\n\nAdapun Gringsing secara bentuk menyimbolkan embrio (awal kehdupan). Tetapi secara kata (Jawa = Jarwa Sodhok) mempunyai arti Gring dan Sing (Jawa gring=sakit, dan Sing = Yang). Jadi Gringsing mempunyai makna jangan pernah sakit & sehat selalu.\n\n\nInformasi Tambahan\nBahan pewarna : Alam\nKondisi kain       : Baik & utuh\n\nHasil kurasi selengkapnya akan diinfokan pada saat live lelang berlangsung.",
  "descEn": "This motif is a development motif that is inside the palace walls (Keraton).\n\nThe Gringsing form symbolizes the embryo (beginning of life). But in the form of word (Java = Jarwa Sodhok) means Gring and Sing (Javanese Gring = sick, and Sing = Yang). So Gringsing means never sick & always healthy.\n\n\nAdditional information\nDye Material: Combination/Natural/Chemical\nFabric condition: Good & Intact\n\nThe full curation results will be announced during the live auction.",
  "priceRange": [3100000, 9999999],
  "size": [245, 104],
  "tipe": "Katun import",
  "eventID": 4,
  "asalDaerah": "Yogyakarta",
  "dyeType": "Pewarna Alam",
  "hargaAwal": 4030000,
  "urlThumbnail": [
    "https://i.ibb.co/6NG71C9/14-1.webp",
    "https://i.ibb.co/82Cbghr/14-2.webp",
    "https://i.ibb.co/XypBQjP/14-3.webp",
    "https://i.ibb.co/ts06CNX/14-4.webp",
    "https://i.ibb.co/z8fZXZF/14-5.webp"
  ],
  "isAvailable": true
}

```

# Event Struct

## Changes:

- ada createdAt,updatedAt, deletedAt tapi gakepake

## Prev:

```
{
  "id": "event-YJ49km_aLlMFpW6t",
  "date": "2022-04-23T02:20:33Z",
  "descId": "Dapatkan batik kuno eksklusif, langka, dan bermakna yang telah terkurasi dan dijamin 100% asli. Mari turut serta berbagi keberkahan kepada para perajin batik lokal dengan memiliki batik yang tersedia di Narauction, karena sebagian dari dana anda akan didonasikan kepada mereka.",
  "descEn": "Get exclusive, rare, and meaningful ancient batik that has been curated and guaranteed to be 100% authentic. Let's participate in sharing the blessings to local batik artisans by having batik available at Narauction because some of your funds will be donated to them.",
  "name": "The Ancient Collection Vol. 2",
  "openingId": "kosong",
  "openingEn": "kosong",
  "foto": ["https://i.ibb.co/D1DyzNT/tac2-websiteheader.png"]
}

```

## New:

```
{
  "id": 4,
  "createdAt": "2022-04-16T12:08:26.267384+07:00",
  "updatedAt": "2022-04-16T12:08:30.920187+07:00",
  "deletedAt": null,
  "name": "The Ancient Collection Vol. 2",
  "date": "2022-04-23T09:20:33+07:00",
  "descId": "Dapatkan batik kuno eksklusif, langka, dan bermakna yang telah terkurasi dan dijamin 100% asli. Mari turut serta berbagi keberkahan kepada para perajin batik lokal dengan memiliki batik yang tersedia di Narauction, karena sebagian dari dana anda akan didonasikan kepada mereka.",
  "descEn": "Get exclusive, rare, and meaningful ancient batik that has been curated and guaranteed to be 100% authentic. Let's participate in sharing the blessings to local batik artisans by having batik available at Narauction because some of your funds will be donated to them.",
  "openingId": "kosong",
  "openingEn": "kosong",
  "foto": ["https://i.ibb.co/D1DyzNT/tac2-websiteheader.png"]
}

```

# UpcomingEvent

## Changes:

- struct event jadi didalem field array "event" (event.id, event.date,...)
- ada createdAt,updatedAt, deletedAt tapi gakepake

## Prev:

```
{
  "id": "event-YJ49km_aLlMFpW6t",
  "date": "2022-04-23T02:20:33Z",
  "descId": "Dapatkan batik kuno eksklusif, langka, dan bermakna yang telah terkurasi dan dijamin 100% asli. Mari turut serta berbagi keberkahan kepada para perajin batik lokal dengan memiliki batik yang tersedia di Narauction, karena sebagian dari dana anda akan didonasikan kepada mereka.",
  "descEn": "Get exclusive, rare, and meaningful ancient batik that has been curated and guaranteed to be 100% authentic. Let's participate in sharing the blessings to local batik artisans by having batik available at Narauction because some of your funds will be donated to them.",
  "name": "The Ancient Collection Vol. 2",
  "openingId": "kosong",
  "openingEn": "kosong",
  "foto": ["https://i.ibb.co/D1DyzNT/tac2-websiteheader.png"],
  "dayDelta": 6,
  "eventStatus": "upcoming event",
  "itemCount": 10,
  "fotoItem": [
    "https://i.ibb.co/rQQKTdt/14-1.jpg",
    "https://i.ibb.co/DGwV8Mr/07-1.jpg",
    "https://i.ibb.co/hBsfGsJ/12-1.jpg",
    "https://i.ibb.co/CPxC43M/b-01-1.jpg",
    "https://i.ibb.co/nMSnfJR/b-02-1.jpg"
  ]
}

```

## New:

```
{
  "event": {
    "id": 4,
    "createdAt": "2022-04-16T12:08:26.267384+07:00",
    "updatedAt": "2022-04-16T12:08:30.920187+07:00",
    "deletedAt": null,
    "name": "The Ancient Collection Vol. 2",
    "date": "2022-04-23T09:20:33+07:00",
    "descId": "Dapatkan batik kuno eksklusif, langka, dan bermakna yang telah terkurasi dan dijamin 100% asli. Mari turut serta berbagi keberkahan kepada para perajin batik lokal dengan memiliki batik yang tersedia di Narauction, karena sebagian dari dana anda akan didonasikan kepada mereka.",
    "descEn": "Get exclusive, rare, and meaningful ancient batik that has been curated and guaranteed to be 100% authentic. Let's participate in sharing the blessings to local batik artisans by having batik available at Narauction because some of your funds will be donated to them.",
    "openingId": "kosong",
    "openingEn": "kosong",
    "foto": ["https://i.ibb.co/D1DyzNT/tac2-websiteheader.png"]
  },
  "dayDelta": 6,
  "status": "upcoming event",
  "itemCount": 10,
  "fotoItem": [
    "https://i.ibb.co/rQQKTdt/14-1.jpg",
    "https://i.ibb.co/DGwV8Mr/07-1.jpg",
    "https://i.ibb.co/hBsfGsJ/12-1.jpg",
    "https://i.ibb.co/CPxC43M/b-01-1.jpg",
    "https://i.ibb.co/nMSnfJR/b-02-1.jpg"
  ]
}

```
