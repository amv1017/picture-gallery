package database

import (
	"strconv"
	"github.com/amv1017/picture-gallery/models"
	"gorm.io/gorm"
)

var (

	DB *gorm.DB

	sp = []models.Painting {

		{ // 0
			Title: "The Rooks Have Come Back", 
			Description: "The Rooks Have Come Back (1871) is considered by many critics to be the high point in Savrasov’s artistic career. Using a common, even trivial, episode of birds returning home, and an extremely simple landscape, Savrasov emotionally showed the transition of nature from winter to spring. It was a new type of lyrical landscape painting, called later by critics the mood landscape. The painting brought him fame. In 1870, he became a member of the Peredvizhniki group, breaking with government-sponsored academic art.",
			Url: "https://uploads6.wikiart.org/images/aleksey-savrasov/rooks-have-returned-1871(1).jpg!Large.jpg",
			Author: "Aleksey Savrasov",
			Genre: "Landscape",
		},
		{ // 1
			Title: "The Ninth Wave",
			Description: "The Ninth Wave (Russian: Девятый вал, Dyevyatiy val) is an 1850 painting by Russian marine painter Ivan Aivazovsky. It is his best-known work.[1][2] The title refers to an old sailing expression referring to a wave of incredible size that comes after a succession of incrementally larger waves.[3] It depicts a sea after a night storm and people facing death attempting to save themselves by clinging to debris from a wrecked ship. The debris, in the shape of the cross, appears to be a Christian metaphor for salvation from the earthly sin. The painting has warm tones, which reduce the sea's apparent menacing overtones and a chance for the people to survive seems plausible. This painting shows both the destructiveness and beauty of nature.",
			Url: "https://uploads8.wikiart.org/00129/images/ivan-aivazovsky/the-ninth-wave.jpg!Large.jpg",
			Author: "Ivan Aivazovsky",
			Genre: "Landscape",
		},
		{ // 2
			Title: "The Last Day of Pompeii",
			Description: "The Last Day of Pompeii is a large history painting by Karl Bryullov produced in 1830–1833 on the subject of the eruption of Mount Vesuvius in AD 79. It is notable for its positioning between Neoclassicism, the predominant style in Russia at the time, and Romanticism as increasingly practised in France. The painting was received to near universal acclaim and made Bryullov the first Russian painter to have an international reputation. In Russia it was seen as proving that Russian art was as good as art practised in the rest of Europe. Critics in France and Russia both noted, however, that the perfection of the classically modelled bodies seemed to be out of keeping with their desperate plight and the overall theme of the painting, which was a Romantic one of the sublime power of nature to destroy man's creations.",
			Url: "https://uploads4.wikiart.org/images/karl-bryullov/the-last-day-of-pompeii-1833.jpg!Large.jpg", 
			Author: "Karl Bryullov",
			Genre: "Landscape",
		},
		{ // 3
			Title: "Morning in a Pine Forest",
			Description: "The Morning in a Pine Forest (Russian: Утро в сосновом лесу) is a painting by Russian artists Ivan Shishkin and Konstantin Savitsky. The bears were painted by Savitsky,[1] but the art collector Pavel Tretyakov effaced his signature, stating that from idea until performance, everything discloses the painting manner and creative method peculiar just to Shishkin,[2] so the painting is now credited solely to Shishkin.",
			Url: "https://uploads5.wikiart.org/images/ivan-shishkin/morning-in-a-pine-forest-1889.jpg!Large.jpg",
			Author: "Ivan Shishkin",
			Genre: "Landscape",
		},
		{ // 4
			Title: "The Field of Wheat",
			Description: "Rye (1878) is one of this artist's best canvasses. By embellishing it, Shishkin transfigured the actual landscape that inspired the painting. Nevertheless, the painter remained faithful to reality and accurately depicted the details of the landscape: the flowers, trees and grain. The proportionate enlargement of details and the combination of various subjects - of the field and the forest, which in reality rarely coexist - create an epic image of an abundant Russian nature.",
			Url: "https://uploads8.wikiart.org/images/ivan-shishkin/the-field-of-wheat-1878.jpg!Large.jpg",
			Author: "Ivan Shishkin",
			Genre: "Landscape",
		},
		{ // 5
			Title: "Golden Autumn",
			Description: "Many features of the Golden Autumn illustrate the increased interest in the problematics of form and the issues of the quality of painting and style in the 1890s. This work combines the immediacy of plein-air observations with the search for an updated plastic solution. A completely traditional painting style mixes here with a free, almost impressionistic interpretation of individual details. At the same time, colour does not dissolve in light, as in classical Impressionism, but retains its intensity, which, together with expressive pastose brushwork, enhances the emotional emphasis of the canvas.",
			Url: "https://uploads3.wikiart.org/images/isaac-levitan/golden-autumn-1895.jpg!Large.jpg",
			Author: "Isaac Levitan",
			Genre: "Landscape",
		},
		{ // 6
			Title: "March",
			Description: "March is Levitan's most upbeat work. The landscape reflects unique features of Russian Impressionism. Nature is transformed with bright colours, all of which are imbued with light. The contrasting yellow and blue shades 'sound' sonorously and joyfully. Everything is still covered in snow, but each detail of the landscape is full of spring exultation. Heated by the sun, the house wall seems to radiate a yellow light that flares up everywhere with golden overtones warming up cold colours. The picture's foreground is full of movement: textured dabs convey both the play of sunlight and the looseness of melting snow. Nature seems to be undergoing a 'set change'. The tree trunks stretching towards the sun are arranged symmetrically, like side-scenes. Moved far into the background is a pine forest whose colours faded over the winter. The eye is lost in the abundance of impressions, but the drowsy horse dosing under the sun catches the viewer's attention. The impressionist mobility of the landscape expresses various rhythms of awakening nature.",
			Url: "https://uploads4.wikiart.org/images/isaac-levitan/march-1895.jpg!Large.jpg",
			Author: "Isaac Levitan",
			Genre: "Landscape",
		},
		{ // 7
			Title: "Lake Ladoga",
			Description: "All the summer of 1872, Kuinji spent on Valaam. The result was the emergence of this ingenious creation. The landscape is at a crossroads. On the one hand, this is a realistic understanding by the artist of nature. But at the same time, the characteristic influence of romanticism is still felt. The viewer feels the romantic features in the foreground, which is deliberately highlighted by the light of a thunderstorm. Kuindzhi was able to overcome in his work the over-exertion with which a special state of nature was transmitted, which was characteristic of the works of late romantics. The landscape is written with a special grace. Shades of light are incredibly subtle. You can feel the incredible integrity, which completely removes the contrasts of light.",
			Url: "https://uploads3.wikiart.org/images/arkhip-kuindzhi/lake-ladoga-1873.jpg!Large.jpg",
			Author: "Arkhip Kuindzhi",
			Genre: "Landscape",
		},
		{ // 8
			Title: "Moscow Court",
			Description: "Vasily Polenov often repeated his most successful works. Depicting a typical corner of old Moscow, this is a later version of one of the artist’s most famous paintings — Moscow Courtyard (1878, Tret. Gal.). Combining elements of a landscape and genre scene, the canvas is filled with the poetry of a peaceful existence and the virtues of a lazy existence. Polenov fully reveals his mastery of plein-air painting in this work, making it one of the most important canvases in his oeuvre and in the history of Russian landscape painting in general. The artist works in light colours, softening the reflections and avoiding dark tones and bright patches. Permeated with light and air, this picture is a symbol of old Russia.",
			Url: "https://uploads3.wikiart.org/images/vasily-polenov/moscow-court-1877.jpg!Large.jpg",
			Author: "Vasily Polenov",
			Genre: "Landscape",
		},



		{ // 9
			Title: "Rider",
			Description: "The depicted are sisters Giovannina and Amazilia Pacini, the foster daughters of Countess Y.P. Samoilova, with whom Bryullov had a close friendship.What is important in the painting The Rider is not so much the story as the action. The older of the two sisters abruptly stops her excited horse, though she remains absolutely calm. Wild force submitting to fragile beauty is one of the favourite themes of Romanticism. The girl’s face is perfectly beautiful. The Italian type of appearance was considered to be perfect in Bryullov’s day and the artist was delighted to use this to advantage. The refined play of colours, the sparkling fabrics – every detail literally speaks of the grandeur of this best of all possible worlds.",
			Url: "https://uploads7.wikiart.org/images/karl-bryullov/rider-portrait-of-giovanina-and-amacilia-pacini-the-foster-children-of-countess-yu-p-samoilova.jpg!Large.jpg",
			Author: "Karl Bryullov",
			Genre: "Portrait",
		},
		{ // 10
			Title: "Portrait of M.I.Lopukhina",
			Description: "In the portrait of Maria Ivanovna Lopukhina (1779–1803), née Countess Tolstaya, we see an embodiment of the aesthetic ideal of Sentimentalism. The artist was especially attracted to the nuances in human condition and the life of the soul. An elegiac dreaminess and languorous tenderness permeates the entire artistic fabric of the work. The vague, blurry contours, complex gradations of cool colours with iridescent blue, lilac, greenish, silver notes, join the figure with the landscape into a single harmonious and musical image. The bright blue of the cornflowers flowering amidst the rye reflects the intense tone of the model’s belt. The golden ears repeat the colour of the chain decorating Lopukhina’s arm. The rose and lilac scarf matches up with the drooping roses. Cornflowers, ears of rye and the branches bending down all embody a country landscape corresponding to the delicately simple clothing of the model and the mild, thoughtful expression on her face. Man in the age of Sentimentalism strived to blend in with natural beauty and to feel that he was part of living nature.",
			Url: "https://uploads5.wikiart.org/images/vladimir-borovikovsky/portrait-of-m-i-lopukhina-1797.jpg!Large.jpg",
			Author: "Vladimir Borovikovsky",
			Genre: "Portrait",
		},
		{ // 11
			Title: "Portrait of the Russian poet Gavril Derzhavin",
			Description: "Gavriil (Gavrila) Romanovich Derzhavin (Russian: Гаврии́л (Гаври́ла) Рома́нович Держа́вин, IPA: [ɡɐˈvrilə rɐˈmanəvʲɪtɕ dʲɪrˈʐavʲɪn] (About this soundlisten); 14 July 1743 – 20 July 1816) was one of the most highly esteemed Russian poets before Alexander Pushkin, as well as a statesman. Although his works are traditionally considered literary classicism, his best verse is rich with antitheses and conflicting sounds in a way reminiscent of John Donne and other metaphysical poets.",
			Url: "https://uploads7.wikiart.org/images/vladimir-borovikovsky/portrait-of-the-russian-poet-gavril-derzhavin-1795.jpg!Large.jpg",
			Author: "Vladimir Borovikovsky",
			Genre: "Portrait",
		},
		{ // 12
			Title: "Portrait of Unknown Woman",
			Description: "The viewer is intrigued both by the heroine of the painting and its name. The artist depicted a young woman in a carriage against the Anichkov Palace in St Petersburg. The woman is not so much beautiful as she is impressive and 'chic'. Her costume corresponds to the latest fashion of the time and indicates that she belongs to 'ladies of the demi-monde'. It’s not without reason that critics used to call her 'the courtesan in a carriage', 'the dear camellia' and 'an offspring of big cities'. Kramskoi emphasises certain demonism in his heroine’s features - sensuous lips, hazy eyes, thick curved eyebrows. The topic of the beauty of sin would become popular among the next generation of Russian artists. The painting is done in an unusually bright, sappy, relaxed manner. Kramskoi clearly intended to show off his outstanding painting mastery.",
			Url: "https://uploads0.wikiart.org/00200/images/ivan-kramskoy/kramskoy-portrait-of-a-woman.jpg!Large.jpg",
			Author: "Ivan Kramskoy",
			Genre: "Portrait",
		},
		{ // 13
			Title: "Portrait of the Author Feodor Dostoyevsky",
			Description: "Fyodor Mikhailovich Dostoyevsky (1821–1881) is a great writer, the most widely read Russian author abroad, creator of the novels Crime and Punishment (1866), The Idiot (1868), Demons (1871–1872), The Adolescent (1875) and The Brothers Karamazov (1879–1880); from 1873 to 1881 – the publisher of A Writer's Diary. The portrait was made when Dostoyevsky was working on Demons. The painting's psychological dramatic effect lies in the contrast between the writer's concentrated, frozen face ('the head with its frozen suffering' – as I.N. Kramskoi put it) and his tightly folded arms that seem to remember being in shackles. Dostoyevsky was sentenced to death for his part in the revolutionary circle of M.V. Butashevich-Petrashevsky, but at the last moment his punishment was changed to hard labour. The writer's spouse A.G. Dostoyevskaya said that 'Perov captured... Dostoyevsky's 'creative moment' ... he seems to be 'peering into himself'.' Dostoyevsky is depicted in a pose that is similar to that of Christ in Kramskoi's Christ in the Desert. This resemblance was not coincidental for the writer's contemporaries. According to Kramskoi, 'the main virtue [of the portrait] is obviously the fact that it expressed the character of the famous writer and man.",
			Url: "https://uploads7.wikiart.org/00341/images/vasily-perov/portrait-of-dostoyevsky.jpg!Large.jpg",
			Author: "Vasily Perov",
			Genre: "Portrait",
		},
		{ // 14
			Title: "Portrait of the Author Ivan Turgenev",
			Description: "Ivan Sergeyevich Turgenev was a Russian novelist, short story writer, poet, playwright, translator and popularizer of Russian literature in the West. His first major publication, a short story collection entitled A Sportsman's Sketches (1852), was a milestone of Russian realism. His novel Fathers and Sons (1862) is regarded as one of the major works of 19th-century fiction.",
			Url: "https://uploads7.wikiart.org/images/vasily-perov/portrait-of-the-author-ivan-turgenev-1872.jpg!Large.jpg",
			Author: "Vasily Perov",
			Genre: "Portrait",
		},
		{ // 15
			Title: "Portrait of the Composer Modest Musorgsky",
			Description: "Modest Petrovich Mussorgsky (1839–1881) is a Russian composer. His main works are the operas Boris Godunov (1872), Khovanshchina (completed by N.A. Rimsky-Korsakov, 1883), The Fair at Sorochyntsi (completed by C.A. Cui, 1916) and the piano suite Pictures at an Exhibition (1874). Mussorgsky's portrait is a masterpiece in Repin's art and in the entire Russian portraiture of the 1880s. It was painted 10 days before the composer's death, in the hospital, in four sessions, and, according to an eye-witness, 'with every possible inconvenience; the painter did not even have an easel and he had to perch somehow near the desk at which Mussorgsky was sitting in a hospital armchair'. The artist depicts the terminally ill man with rare delicacy. The indeterminate background creates an illusion of open space. Musorgsky seems to be depicted against the sky, his gaze is far away. Repin refused to accept the payment for the work donating the money to the composer's monument.",
			Url: "https://uploads1.wikiart.org/images/ilya-repin/portrait-of-the-composer-modest-musorgsky-1881.jpg!Large.jpg",
			Author: "Ilya Repin",
			Genre: "Portrait",
		},



		{ // 16
			Title: "Bird Cherry In A Glass",
			Description: "Kuzma Sergeevich Petrov-Vodkin was a Russian artist who studied at art school in St Petersburg, after which he went to Moscow and Munich. Petrov-Vodkin’s most famous painting Bathing of a Red Horse 1912 forms part of the permanent collection at Moscow’s Tretyakov State Gallery. The Communists denounced his work but in the 1970’s his literary essays and paintings were rediscovered and there is an entire room at the Russian Museum in St Petersburg dedicated to his paintings.",
			Url: "https://uploads0.wikiart.org/images/kuzma-petrov-vodkin/bird-cherry-in-a-glass-1932.jpg!Large.jpg",
			Author: "Kuzma Petrov-Vodkin",
			Genre: "Still Life",
		},
		{ // 17
			Title: "Cornflowers",
			Description: "Levitan is known to everyone as the most famous Russian landscape painter. In this picture, we see an image not quite familiar to him – a still life. Nothing distracts the viewer from the spot of meaning: the artist has not created a single detail around a vase with flowers, for which the viewer could catch the eye. All attention is paid to flowers: divinely beautiful cornflowers in their simplicity. Simple wildflowers amaze with their magnificence. It would seem that ordinary cornflowers, the same, similar to each other. But each of this small bouquet of cornflower flower attracts separate attention. Each of them has its own special color. Here the viewer will find bright blue cornflowers, and with a slight shade of purple, and dove, and pink, and even white cornflowers.",
			Url: "https://uploads1.wikiart.org/images/isaac-levitan/cornflowers-1894.jpg!Large.jpg",
			Author: "Isaac Levitan",
			Genre: "Still Life",
		},
		{ // 18
			Title: "A Bunch of Flowers. Phloxes",
			Description: "Kramskoy – Russian artist, from all genres, preferred genre and historical painting. Sometimes he painted portraits, criticized other people’s paintings, tried to find such measures by which one could accurately determine the value of objects of art. He fiercely defended the idea that the artist should be not just an observer, but a teacher, and that the pictures should serve a greater purpose than just a reflection of reality. In his opinion, they had to develop morality and offer a person to make a moral choice, had to inculcate taste and reflect deeply folk, national subjects, reminding people who they were and where they came from. The pictures in his world were supposed to be a guide, a light in the dark, and this very striving for light, for freedom, for good, can be seen not only in the works of the artist, but also in his life.",
			Url: "https://uploads8.wikiart.org/images/ivan-kramskoy/bouquet-of-flowers-1884.jpg!Large.jpg",
			Author: "Ivan Kramskoy",
			Genre: "Still Life",
		},
		{ // 19
			Title: "Still Life (1911)",
			Description: "Malevich is a famous Russian impressionist, the most famous painting of which 'Black Square' still causes controversy and bewilderment – some see in it a deep deep meaning, others argue with them, speaking of the picture as a meaningless daub. However, before coming to the 'Black Square', symbolizing the end of everything, Malevich spent a long time looking for himself and his style, trying a variety of options. Realism disliked him, classicism seemed boring. He considered senseless to carry life into art, saying that, on the contrary, it should be carried into art, for this is the only way to make it truly beautiful. Therefore, in his quest he turned to impressionism and cubism, but never to realism. He was close to bright colors, strong, harsh emotions, and were not at all near pastoral, tenderness and trembling. He wanted to change his life and wanted to do it through their art.",
			Url: "https://uploads7.wikiart.org/images/kazimir-malevich/still-life.jpg!Large.jpg",
			Author: "Kazimir Malevich",
			Genre: "Still Life",
		},
		{ // 20
			Title: "Apples and Leaves",
			Description: "When this picture was written, Ilya Repin was at the peak of creativity, he created one masterpiece after another, had a fee, he had a wonderful family, with which he moved to the capital, here he teamed up with the most prominent artists of Russia who met him and his works complacently, exhibited with a cycle of masterfully written portraits, sold his work for good money at that time, acquired new friends in Moscow. Of course, all these circumstances could not but affect the mood of the artist, and he created such an unusual work for himself as 'Apples and Leaves'. Why unusual? Because Ilya Efimovich is a virtuoso of historical and mythological themes, portrait painting, psychological, imbued with feelings and emotions of paintings. Still life was not peculiar to this artist.",
			Url: "https://uploads4.wikiart.org/images/ilya-repin/apples-and-leaves-1879.jpg!Large.jpg",
			Author: "Ilya Repin",
			Genre: "Still Life",
		},

	}

	sa = []models.Author { 
		{ Name: "Aleksey Savrasov", Paintings: []models.Painting{ 
				sp[0],
			},
		},
		{ Name: "Ivan Aivazovsky", Paintings: []models.Painting{ 
				sp[1],
			},  
		},
		{ Name: "Karl Bryullov", Paintings: []models.Painting{ 
				sp[2],
				sp[9],
			},  
		},
		{ Name: "Vladimir Borovikovsky", Paintings: []models.Painting{ 
				sp[10],
				sp[11],
			},  
		},
		{ Name: "Kuzma Petrov-Vodkin", Paintings: []models.Painting{ 
				sp[16],
			},  
		},
		{ Name: "Ivan Shishkin", Paintings: []models.Painting{ 
				sp[3],
				sp[4],
			},  
		},
		{ Name: "Isaac Levitan", Paintings: []models.Painting{ 
				sp[5],
				sp[6],
				sp[17],
			},  
		},
		{ Name: "Arkhip Kuindzhi", Paintings: []models.Painting{ 
				sp[7],
			},  
		},
		{ Name: "Vasily Polenov", Paintings: []models.Painting{ 
				sp[8],
			},  
		},
		{ Name: "Ivan Kramskoy", Paintings: []models.Painting{ 
				sp[12],
				sp[18],
			},  
		},
		{ Name: "Vasily Perov", Paintings: []models.Painting{ 
				sp[13],
				sp[14],
			},  
		},
		{ Name: "Ilya Repin", Paintings: []models.Painting{ 
				sp[15],
				sp[20],
			},  
		},
		{ Name: "Kazimir Malevich", Paintings: []models.Painting{ 
				sp[19],
			},  
		},

	}

	sg = []models.Genre {
		{ Sign: "Landscape" },
		{ Sign: "Portrait" },
		{ Sign: "Still Life" },
	}

)

func FillDatabase() {

	for i := range sg {
		DB.Create(&sg[i])
	}

	for i := range sa {
		DB.Create(&sa[i])
	}

	gb := []int{1,2,3,8,9,10,11,13,14}
	for i := range gb {
		DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title,painting_url) VALUES (`+
		strconv.Itoa(int(sg[0].ID))+`,'`+sg[0].Sign+`',`+strconv.Itoa(gb[i])+`,'`+sp[i].Title+`','`+sp[i].Url+`')`)
	}

	gb = []int{4,5,6,15,17,18,19}
	for i := range gb {
		DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title,painting_url) VALUES (`+
		strconv.Itoa(int(sg[1].ID))+`,'`+sg[1].Sign+`',`+strconv.Itoa(gb[i])+`,'`+sp[i+9].Title+`','`+sp[i+9].Url+`')`)
	}

	gb = []int{7,12,16,21,20}
	for i := range gb {
		DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title,painting_url) VALUES (`+
		strconv.Itoa(int(sg[2].ID))+`,'`+sg[2].Sign+`',`+strconv.Itoa(gb[i])+`,'`+sp[i+16].Title+`','`+sp[i+16].Url+`')`)
	}

}
