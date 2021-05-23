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
		},
		{ // 1
			Title: "The Ninth Wave",
			Description: "The Ninth Wave (Russian: Девятый вал, Dyevyatiy val) is an 1850 painting by Russian marine painter Ivan Aivazovsky. It is his best-known work.[1][2] The title refers to an old sailing expression referring to a wave of incredible size that comes after a succession of incrementally larger waves.[3] It depicts a sea after a night storm and people facing death attempting to save themselves by clinging to debris from a wrecked ship. The debris, in the shape of the cross, appears to be a Christian metaphor for salvation from the earthly sin. The painting has warm tones, which reduce the sea's apparent menacing overtones and a chance for the people to survive seems plausible. This painting shows both the destructiveness and beauty of nature.",
			Url: "https://uploads8.wikiart.org/00129/images/ivan-aivazovsky/the-ninth-wave.jpg!Large.jpg",
		},
		{ // 2
			Title: "The Last Day of Pompeii",
			Description: "The Last Day of Pompeii is a large history painting by Karl Bryullov produced in 1830–1833 on the subject of the eruption of Mount Vesuvius in AD 79. It is notable for its positioning between Neoclassicism, the predominant style in Russia at the time, and Romanticism as increasingly practised in France. The painting was received to near universal acclaim and made Bryullov the first Russian painter to have an international reputation. In Russia it was seen as proving that Russian art was as good as art practised in the rest of Europe. Critics in France and Russia both noted, however, that the perfection of the classically modelled bodies seemed to be out of keeping with their desperate plight and the overall theme of the painting, which was a Romantic one of the sublime power of nature to destroy man's creations.",
			Url: "https://uploads4.wikiart.org/images/karl-bryullov/the-last-day-of-pompeii-1833.jpg!Large.jpg", 
		},
		{ // 3
			Title: "Rider",
			Description: "The depicted are sisters Giovannina and Amazilia Pacini, the foster daughters of Countess Y.P. Samoilova, with whom Bryullov had a close friendship.What is important in the painting The Rider is not so much the story as the action. The older of the two sisters abruptly stops her excited horse, though she remains absolutely calm. Wild force submitting to fragile beauty is one of the favourite themes of Romanticism. The girl’s face is perfectly beautiful. The Italian type of appearance was considered to be perfect in Bryullov’s day and the artist was delighted to use this to advantage. The refined play of colours, the sparkling fabrics – every detail literally speaks of the grandeur of this best of all possible worlds.",
			Url: "https://uploads7.wikiart.org/images/karl-bryullov/rider-portrait-of-giovanina-and-amacilia-pacini-the-foster-children-of-countess-yu-p-samoilova.jpg!Large.jpg",
		},
		{ // 4
			Title: "Portrait of M.I.Lopukhina",
			Description: "In the portrait of Maria Ivanovna Lopukhina (1779–1803), née Countess Tolstaya, we see an embodiment of the aesthetic ideal of Sentimentalism. The artist was especially attracted to the nuances in human condition and the life of the soul. An elegiac dreaminess and languorous tenderness permeates the entire artistic fabric of the work. The vague, blurry contours, complex gradations of cool colours with iridescent blue, lilac, greenish, silver notes, join the figure with the landscape into a single harmonious and musical image. The bright blue of the cornflowers flowering amidst the rye reflects the intense tone of the model’s belt. The golden ears repeat the colour of the chain decorating Lopukhina’s arm. The rose and lilac scarf matches up with the drooping roses. Cornflowers, ears of rye and the branches bending down all embody a country landscape corresponding to the delicately simple clothing of the model and the mild, thoughtful expression on her face. Man in the age of Sentimentalism strived to blend in with natural beauty and to feel that he was part of living nature.",
			Url: "https://uploads5.wikiart.org/images/vladimir-borovikovsky/portrait-of-m-i-lopukhina-1797.jpg!Large.jpg",
		},
		{ // 5
			Title: "Portrait of the Russian poet Gavril Derzhavin",
			Description: "Gavriil (Gavrila) Romanovich Derzhavin (Russian: Гаврии́л (Гаври́ла) Рома́нович Держа́вин, IPA: [ɡɐˈvrilə rɐˈmanəvʲɪtɕ dʲɪrˈʐavʲɪn] (About this soundlisten); 14 July 1743 – 20 July 1816) was one of the most highly esteemed Russian poets before Alexander Pushkin, as well as a statesman. Although his works are traditionally considered literary classicism, his best verse is rich with antitheses and conflicting sounds in a way reminiscent of John Donne and other metaphysical poets.",
			Url: "https://uploads7.wikiart.org/images/vladimir-borovikovsky/portrait-of-the-russian-poet-gavril-derzhavin-1795.jpg!Large.jpg",
		},
		{ // 6
			Title: "Bird cherry in a glass",
			Description: "Kuzma Sergeevich Petrov-Vodkin was a Russian artist who studied at art school in St Petersburg, after which he went to Moscow and Munich. Petrov-Vodkin’s most famous painting Bathing of a Red Horse 1912 forms part of the permanent collection at Moscow’s Tretyakov State Gallery. The Communists denounced his work but in the 1970’s his literary essays and paintings were rediscovered and there is an entire room at the Russian Museum in St Petersburg dedicated to his paintings.",
			Url: "https://uploads0.wikiart.org/images/kuzma-petrov-vodkin/bird-cherry-in-a-glass-1932.jpg!Large.jpg",
		},


	}

	setup_author = []models.Author { 
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
				sp[3],
			},  
		},
		{ Name: "Vladimir Borovikovsky", Paintings: []models.Painting{ 
				sp[4],
				sp[5],
			},  
		},
		{ Name: "Kuzma Petrov-Vodkin", Paintings: []models.Painting{ 
				sp[6],
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

	for i := range setup_author {
		DB.Create(&setup_author[i])
	}



	DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title) VALUES (`+strconv.Itoa(int(sg[0].ID))+`,'`+sg[0].Sign+`',`+strconv.Itoa(int(1))+`,'`+sp[0].Title+`')`)
	DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title) VALUES (`+strconv.Itoa(int(sg[0].ID))+`,'`+sg[0].Sign+`',`+strconv.Itoa(int(2))+`,'`+sp[1].Title+`')`)


}

