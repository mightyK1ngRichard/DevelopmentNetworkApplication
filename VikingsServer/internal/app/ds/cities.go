package ds

type City struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Area        string `json:"area"`
}

type CityViewData struct {
	Cities   []City
	LookAlso []City
}

func GetCityViewData() *CityViewData {
	return &CityViewData{
		Cities: []City{
			{ID: 0, Name: "Вальхалла", Description: "Вальхалла - мифический зал, куда отправляются воины погибшие в битвах.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=175428ff6c8f7e34ea3bc9b72ddd99d9128e2ac0-9216025-images-thumbs&n=13"},
			{ID: 1, Name: "Рагнарок", Description: "Рагнарок - конец света в викингской мифологии.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=49f075566277c2cdfd8ecbec770ff8dba800e8c5-9848534-images-thumbs&n=13"},
			{ID: 2, Name: "Норсгард", Description: "Норсгард - крепость викингов на севере.", Area: "1000 кв. м.", ImageURL: "http://localhost:7070/static/img/image1.jpeg"},
			{ID: 3, Name: "Эйрстад", Description: "Эйрстад - город воздуха, летающий на драконах.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image2.jpg"},
			{ID: 4, Name: "Хельхейм", Description: "Хельхейм - царство мертвых в викингской мифологии.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=7c8984c33159f006c9b671cb30c210deb9db10de-9290563-images-thumbs&n=13"},
			{ID: 5, Name: "Фьордхейм", Description: "Фьордхейм - город на берегу фьорда.", Area: "1500 кв. м.", ImageURL: "http://localhost:7070/static/img/image3.jpg"},
			{ID: 6, Name: "Винтерфелл", Description: "Винтерфелл - зимний город с крепостью.", Area: "2000 кв. м.", ImageURL: "http://localhost:7070/static/img/image4.jpg"},
			{ID: 7, Name: "Мидгард", Description: "Мидгард - мир людей в скандинавской мифологии.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image5.jpg"},
			{ID: 8, Name: "Йотунхейм", Description: "Йотунхейм - царство великанов в викингской мифологии.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image6.jpg"},
			{ID: 9, Name: "Асгард", Description: "Асгард - дом богов в скандинавской мифологии.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image7.jpg"},
			{ID: 10, Name: "Улфхейм", Description: "Улфхейм - город викингов с волками в гербе.", Area: "1200 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=99b86aff2164bdd6314ca0e5c7a88083376ebc62-9065868-images-thumbs&n=13"},
			{ID: 11, Name: "Фростгард", Description: "Фростгард - город, окруженный вечной метелью.", Area: "800 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=5cad3917adaf121cc7bd04d42248f857bbd5d29a-9185064-images-thumbs&n=13"},
			{ID: 12, Name: "Драккенсберг", Description: "Драккенсберг - город, славящийся драконами.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=4d37154e953fbff590c883a49fa3ed668f8ac58a-8491894-images-thumbs&n=13"},
			{ID: 13, Name: "Стормхейвен", Description: "Стормхейвен - город на берегу бурного моря.", Area: "1800 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=c38b0ff5c2b07cb3a81a6f0dfadb21598d885b82-9847423-images-thumbs&n=13"},
			{ID: 14, Name: "Гундархольм", Description: "Гундархольм - остров викингов.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=cb87b162cf9776257cbbad336aec6a17a8725dd4-9072018-images-thumbs&n=13"},
			{ID: 15, Name: "Скади", Description: "Скади - горный поселок на высоких скалах.", Area: "2500 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=8e0ab64ed97fe057032be5ebc0398ab3-4568063-images-thumbs&n=13"},
			{ID: 16, Name: "Локисберг", Description: "Локисберг - город, почитающий бога Локи.", Area: "1600 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=49f075566277c2cdfd8ecbec770ff8dba800e8c5-9848534-images-thumbs&n=13"},
			{ID: 17, Name: "Хельгринд", Description: "Хельгринд - город, связанный с богиней Хель.", Area: "1400 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=939d8637492973bae36925c4c9bae5293f474b23-10878189-images-thumbs&n=13"},
			{ID: 18, Name: "Гримсберг", Description: "Гримсберг - город, где проживают ведьмы.", Area: "1100 кв. м.", ImageURL: "https://avatars.mds.yandex.net/i?id=f11cf4cbc3aeabdb48d596f111b5b96fddd937c2-8341813-images-thumbs&n=13"},
			{ID: 19, Name: "Гунлейф", Description: "Гунлейф - город-корабль викингов.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=644eb1df0fb871b49ad3ab951bd2bfa31a9a5f3e-9094594-images-thumbs&n=13"},
		},
	}
}
