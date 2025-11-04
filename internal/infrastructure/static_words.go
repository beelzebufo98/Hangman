package infrastructure

import (
	"crypto/rand"
	"math/big"
	"sort"
	"strings"

	"github.com/beelzebufo98/Hangman/internal/domain"
)

type StaticWordProvider struct {
	data map[string]map[domain.Difficulty][]domain.Word
}

func NewStaticWordProvider() *StaticWordProvider {
	return &StaticWordProvider{data: wordsData()}
}

func (p *StaticWordProvider) GetCategories() []string {
	out := make([]string, 0, len(p.data))
	for k := range p.data {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (p *StaticWordProvider) GetLevels(category string) []domain.Difficulty {
	for k := range p.data {
		if strings.EqualFold(k, category) {
			levels := make([]domain.Difficulty, 0, len(p.data[k]))
			for lv := range p.data[k] {
				levels = append(levels, lv)
			}
			sort.Slice(levels, func(i, j int) bool { return levels[i] < levels[j] })
			return levels
		}
	}
	return []domain.Difficulty{domain.Easy, domain.Medium, domain.Hard}
}

func (p *StaticWordProvider) GetRandomWord(category string, level domain.Difficulty) domain.Word {
	var cat string
	for k := range p.data {
		if strings.EqualFold(k, category) {
			cat = k
			break
		}
	}
	if cat == "" {
		for k := range p.data {
			cat = k
			break
		}
	}
	pool := p.data[cat][level]
	if len(pool) == 0 {
		pool = p.data[cat][domain.Easy]
	}
	return pool[randInt(len(pool))]
}

func randInt(n int) int {
	if n <= 0 {
		return 0
	}
	m, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	return int(m.Int64())
}

func wordsData() map[string]map[domain.Difficulty][]domain.Word {
	return map[string]map[domain.Difficulty][]domain.Word{
		"Животные": {
			domain.Easy: {
				{Text: "кот", Hint: "Домашнее животное, ловит мышей"},
				{Text: "собака", Hint: "Лучший друг человека"},
				{Text: "лев", Hint: "Хищник, царь зверей"},
				{Text: "сова", Hint: "Ночная птица с большими глазами"},
				{Text: "заяц", Hint: "Быстро бегает, любит морковку"},
			},
			domain.Medium: {
				{Text: "дельфин", Hint: "Умное морское млекопитающее"},
				{Text: "енот", Hint: "Животное-полоскун с маской на морде"},
				{Text: "фламинго", Hint: "Розовая птица с длинными ногами"},
				{Text: "выдра", Hint: "Живёт у воды, отлично плавает"},
				{Text: "страус", Hint: "Самая большая птица, не умеет летать"},
			},
			domain.Hard: {
				{Text: "хамелеон", Hint: "Рептилия, меняет цвет кожи"},
				{Text: "броненосец", Hint: "Млекопитающее с костяным панцирем"},
				{Text: "бегемот", Hint: "Крупное животное, живёт у воды"},
				{Text: "медоед", Hint: "Смелый зверёк, не боится даже львов"},
				{Text: "утконос", Hint: "Редкое животное с клювом как у птицы"},
			},
		},
		"Фрукты": {
			domain.Easy: {
				{Text: "яблоко", Hint: "Красный, зелёный или жёлтый фрукт"},
				{Text: "груша", Hint: "Сладкий фрукт в форме лампочки"},
				{Text: "слива", Hint: "Фиолетовый или жёлтый плод"},
				{Text: "банан", Hint: "Жёлтый фрукт, любимый обезьянами"},
				{Text: "вишня", Hint: "Небольшая красная ягода на дереве"},
			},
			domain.Medium: {
				{Text: "апельсин", Hint: "Цитрус, богат витамином C"},
				{Text: "персик", Hint: "Фрукт с бархатистой кожурой"},
				{Text: "манго", Hint: "Экзотический плод из тропиков"},
				{Text: "ананас", Hint: "Тропический фрукт с колючей кожурой"},
			},
			domain.Hard: {
				{Text: "гранат", Hint: "Красный плод с множеством зёрен"},
				{Text: "авокадо", Hint: "Фрукт с большой косточкой, часто в салатах"},
				{Text: "киви", Hint: "Коричневая кожура, зелёная мякоть"},
				{Text: "личи", Hint: "Китайский фрукт с белой сладкой мякотью"},
				{Text: "папайя", Hint: "Тропический фрукт оранжевого цвета"},
			},
		},
		"Страны": {
			domain.Easy: {
				{Text: "россия", Hint: "Самая большая страна мира"},
				{Text: "китай", Hint: "Страна Великой стены и драконов"},
				{Text: "египет", Hint: "Страна пирамид и фараонов"},
				{Text: "япония", Hint: "Островное государство восходящего солнца"},
				{Text: "индия", Hint: "Страна Тадж-Махала и йоги"},
			},
			domain.Medium: {
				{Text: "франция", Hint: "Страна Эйфелевой башни и круассанов"},
				{Text: "италия", Hint: "Родина пиццы и Колизея"},
				{Text: "бразилия", Hint: "Футбол, самба и Амазонка"},
				{Text: "германия", Hint: "Страна Бранденбургских ворот"},
				{Text: "испания", Hint: "Фламенко и коррида"},
			},
			domain.Hard: {
				{Text: "аргентина", Hint: "Страна танго и футбола"},
				{Text: "исландия", Hint: "Остров гейзеров и вулканов"},
				{Text: "австралия", Hint: "Земля кенгуру и коал"},
				{Text: "марокко", Hint: "Страна базаров и пустыни Сахара"},
				{Text: "чили", Hint: "Узкая страна вдоль Анд"},
			},
		},
	}
}
