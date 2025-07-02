package main

import (
	"fmt"
	"os"
)

// logResults enregistre les résultats du quiz dans un fichier log.txt
// Il prend en paramètre le quiz, les réponses de l'utilisateur et le score final.
// Chaque question, la réponse donnée par l'utilisateur, la bonne réponse et le résultat (Vrai/Faux) sont enregistrés dans le fichier.
// À la fin, le score final est également enregistré dans le fichier log.txt.
// Si le fichier n'existe pas, il est créé. Si une erreur survient lors de l'ouverture du fichier, un message d'erreur est affiché.
// Cette fonction est utile pour garder une trace des performances de l'utilisateur dans le quiz.

func logResults(quiz []Question, userAnswers []int, score int) {
	// Réinitialisation du fichier log.txt à chaque exécution
	if err := os.WriteFile("log.txt", []byte{}, 0644); err != nil {
		fmt.Println("Erreur lors de la réinitialisation du fichier log:", err)
		return
	}
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier log:", err)
		return
	}
	defer f.Close()

	for i, q := range quiz {
		userAns := "Aucune"
		if i < len(userAnswers) && userAnswers[i] >= 1 && userAnswers[i] <= len(q.Choix) {
			userAns = q.Choix[userAnswers[i]-1]
		}
		result := "Faux"
		if userAns == q.Answer {
			result = "Vrai"
		}
		f.WriteString(fmt.Sprintf("Q%d: %s\nRéponse donnée: %s\nBonne réponse: %s\nRésultat: %s\n\n", i+1, q.Question, userAns, q.Answer, result))
	}
	f.WriteString(fmt.Sprintf("Score final: %d/%d\n\n", score, len(quiz)*10))
}

type Question struct {
	Question string
	Answer   string
	Choix    []string
}

func main() {
	// début de la fonction main
	// Cette fonction est le point d'entrée de l'application Go.

	// Déclaration des variables pour la réponse utilisateur et le score
	var userAnswer int    // Initialisation de la réponse utilisateur à 0
	var score int         // Initialisation du score à 0
	var userAnswers []int // Pour enregistrer les réponses de l'utilisateur

	quiz := []Question{
		{
			Question: "Quelle est la capitale de la France ?",
			Choix:    []string{"Berlin", "Madrid", "Paris", "Rome"},
			Answer:   "Paris",
		},
		{
			Question: "Quelle est la plus grande planète du système solaire ?",
			Choix:    []string{"Terre", "Mars", "Jupiter", "Saturne"},
			Answer:   "Jupiter",
		},
		{
			Question: "Quel est le plus grand océan de la Terre ?",
			Choix:    []string{"Atlantique", "Indien", "Arctique", "Pacifique"},
			Answer:   "Pacifique",
		},
		{
			Question: "Quel est le plus haut sommet du monde ?",
			Choix:    []string{"K2", "Mont Blanc", "Mont Everest", "Kilimandjaro"},
			Answer:   "Mont Everest",
		},
		{
			Question: "Quel est l'élément chimique dont le symbole est 'O' ?",
			Choix:    []string{"Oxygène", "Or", "Osmium", "Oganesson"},
			Answer:   "Oxygène",
		},
		{
			Question: "Quel est le plus grand mammifère terrestre ?",
			Choix:    []string{"Éléphant", "Girafe", "Rhinocéros", "Hippopotame"},
			Answer:   "Éléphant",
		},
		{
			Question: "Quel est le pays le plus peuplé du monde ?",
			Choix:    []string{"Inde", "États-Unis", "Chine", "Indonésie"},
			Answer:   "Inde",

			// Ajoutez en haut d'autres questions si nécessaire
		},
	}

	for i, q := range quiz {
		// Affichage de la question et des choix
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for j, choix := range q.Choix {
			fmt.Printf("%d: %s\n", j+1, choix)
		}
		// Demande la réponse à l'utilisateur
		fmt.Print("Votre réponse (1-4): ")
		fmt.Scan(&userAnswer)
		userAnswers = append(userAnswers, userAnswer)

		// Vérification de la réponse
		if userAnswer >= 1 && userAnswer <= len(q.Choix) && q.Choix[userAnswer-1] == q.Answer {
			fmt.Println("Bonne réponse !")
			score += 10 // Incrémentation du score de 10 points pour une bonne réponse
		} else {
			fmt.Println("Mauvaise réponse !")
		}
	}
	// Affichage du score final et du pourcentage de réussite
	fmt.Printf("Votre score final est : %d/%d\n", score, len(quiz)*10)
	percent := float64(score) / float64(len(quiz)*10) * 100
	fmt.Printf("Pourcentage de réussite : %.2f%%\n", percent)
	// Enregistrement des résultats dans le fichier log.txt
	logResults(quiz, userAnswers, score)
}

// fin de la fonction main
// Cette ligne marque la fin de la fonction main et donc de l'exécution du programme.
