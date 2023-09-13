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
			{ID: 0, Name: "Категат (Хедебю)", Description: "Категат был одним из самых важных викингских городов, расположенных на острове Хейланд в Дании. Город был известен своими морскими торговыми маршрутами и фортификациями.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=175428ff6c8f7e34ea3bc9b72ddd99d9128e2ac0-9216025-images-thumbs&n=13"},
			{ID: 1, Name: "Йорк (Йоркшир, Великобритания)", Description: "Этот викингский город, известный как Йорк, был важным торговым и административным центром во времена викингов.\nДублин (Ирландия): Викинги основали Дублин в 9 веке. Этот город был известен своей торговлей и культурным влиянием в регионе.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=49f075566277c2cdfd8ecbec770ff8dba800e8c5-9848534-images-thumbs&n=13"},
			{ID: 2, Name: "Новгород (Россия)", Description: "Викинги создали поселение в Новгороде, что сделало его важным торговым центром на востоке.", Area: "1000 кв. м.", ImageURL: "http://localhost:7070/static/img/image1.jpeg"},
			{ID: 3, Name: "Лунд (Швеция)", Description: "Лунд был одним из первых викингских городов и центром вероисповедания викингов.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image2.jpg"},
			{ID: 4, Name: "Гриндавик (Исландия)", Description: "Этот город на Исландии служил базой для викингских мореплавателей и рыбаков.", Area: "Неизвестно", ImageURL: "https://avatars.mds.yandex.net/i?id=7c8984c33159f006c9b671cb30c210deb9db10de-9290563-images-thumbs&n=13"},
			{ID: 5, Name: "Висбю (Швеция)", Description: "Этот город на острове Готланд был важным торговым и административным центром викингов.", Area: "1500 кв. м.", ImageURL: "http://localhost:7070/static/img/image3.jpg"},
			{ID: 6, Name: "Лимфьорд (Дания)", Description: "Лимфьорд был важным морским перекрестком и базой для викингских экспедиций.", Area: "2000 кв. м.", ImageURL: "http://localhost:7070/static/img/image4.jpg"},
			{ID: 7, Name: "Рейкьявик (Исландия)", Description: "Викинги основали Рейкьявик, который со временем стал столицей Исландии.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image5.jpg"},
			{ID: 8, Name: "Торшавн (Фарерские острова)", Description: "Этот город служил базой для викингов на Фарерских островах и был важным для контроля торговых путей в Северном Атлантическом регионе.", Area: "Неизвестно", ImageURL: "http://localhost:7070/static/img/image6.jpg"},
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
