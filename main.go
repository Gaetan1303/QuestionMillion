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
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier log:", err)
		return
	}
	// Enregistrement des résultats dans le fichier log.txt (écrasement à chaque exécution)
	if err := os.WriteFile("log.txt", []byte{}, 0644); err != nil {
		fmt.Println("Erreur lors de la réinitialisation du fichier log:", err)
	}
	defer f.Close()

	for i, q := range quiz {
		userAns := "Aucune"
		if i < len(userAnswers) && userAnswers[i] >= 1 && userAnswers[i] <= len(q.Options) {
			userAns = q.Options[userAnswers[i]-1]
		}
		result := "Faux"
		if userAns == q.Answer {
			result = "Vrai"
		}
		f.WriteString(fmt.Sprintf("Q%d: %s\nRéponse donnée: %s\nBonne réponse: %s\nRésultat: %s\n\n", i+1, q.Question, userAns, q.Answer, result))
	}
	f.WriteString(fmt.Sprintf("Score final: %d/%d\n\n", score, len(quiz)*10))
}

// Déclaration de la structure Question pour représenter une question du quiz
type Question struct {
	Question string
	Answer   string
	Options  []string
}

func main() {
	// début de la fonction main
	// Cette fonction est le point d'entrée de l'application Go.

	// Déclaration des variables pour la réponse utilisateur et le score
	var userAnswer int    // Initialisation de la réponse utilisateur à 0
	var score int         // Initialisation du score à 0
	var userAnswers []int // Pour enregistrer les réponses de l'utilisateur

	// Création d'un quiz avec une liste de questions
	quiz := []Question{
		{
			Question: "Quelle est la capitale de la France ?",
			Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
			Answer:   "Paris",
		},
		{
			Question: "Quelle est la plus grande planète du système solaire ?",
			Options:  []string{"Terre", "Mars", "Jupiter", "Saturne"},
			Answer:   "Jupiter",
		},
		{
			Question: "Quel est le plus grand océan de la Terre ?",
			Options:  []string{"Atlantique", "Indien", "Arctique", "Pacifique"},
			Answer:   "Pacifique",
		},
		{
			Question: "Quel est le plus haut sommet du monde ?",
			Options:  []string{"K2", "Mont Blanc", "Mont Everest", "Kilimandjaro"},
			Answer:   "Mont Everest",
		},
		{
			Question: "Quel est l'élément chimique dont le symbole est 'O' ?",
			Options:  []string{"Oxygène", "Or", "Osmium", "Oganesson"},
			Answer:   "Oxygène",
		},
		{
			Question: "Quel est le plus grand mammifère terrestre ?",
			Options:  []string{"Éléphant", "Girafe", "Rhinocéros", "Hippopotame"},
			Answer:   "Éléphant",
		},
		{
			Question: "Quel est le pays le plus peuplé du monde ?",
			Options:  []string{"Inde", "États-Unis", "Chine", "Indonésie"},
			Answer:   "Inde",
			
			// Ajoutez en haut d'autres questions si nécessaire
		}}

	// Boucle sur chaque question du quiz
	for i, q := range quiz {
		// Affichage de la question et des options
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d: %s\n", j+1, option)
		}
		// Demande la réponse à l'utilisateur
		fmt.Print("Votre réponse (1-4): ")
		fmt.Scan(&userAnswer)
		userAnswers = append(userAnswers, userAnswer)

		// Vérification de la réponse
		if userAnswer >= 1 && userAnswer <= len(q.Options) && q.Options[userAnswer-1] == q.Answer {
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
