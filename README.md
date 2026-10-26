# TP-Go-Reloaded

Un outil d'autocorrection de texte en Go : un programme qui lit un fichier, applique toutes les règles de transformation, puis écrit le résultat dans un second fichier.

# Fonctionnalités

- UP
- LOW
- CAP
- HEX
- BIN
- PONCTUATION
- APOSTROPHES
- A --> AN
- TEST

# Technologies utilisées

- **Langage :** Golang. Go est un langage de programmation compilé et concurrent, inspiré de C et de Pascal. Il a été développé par Google à partir d'un concept initial de Robert Griesemer, Rob Pike et Ken Thompson.

- **Outils :** VS Code, GitHub, Google, YouTube et la documentation officielle de Go.

# Le dépôt

git clone https://github.com/li0r-ynov/Application-CLI-Parsing-Go 

# Installation

Étapes :

- Installer VS Code sur votre PC.
- Installer Git.
- Installer le langage de programmation Go.
- Copier l'adresse du dépôt GitHub : https://github.com/li0r-ynov/Application-CLI-Parsing-Go.git
- Ouvrir l'Explorateur de fichiers.
- Sélectionner un dossier dans lequel placer le projet Go-Reloaded.
- Aller dans la barre d'adresse et écrire `cmd`.
- Écrire la commande suivante : git clone https://github.com/li0r-ynov/Application-CLI-Parsing-Go.git
- Retourner dans votre dossier et vérifier que le nouveau dossier est apparu.
- Ouvrir VS Code et aller dans l'onglet « File ».
- Cliquer sur « Open Folder ».
- Sélectionner le dossier que vous venez de créer.

# Utilisation

Commande pour lancer l'exécution du correcteur : go run . sample.txt result.txt

# Fonctions

## 1. Les fonctions de transformation

### UP : Mettre en majuscules

**Rôle :** Transforme un mot ou plusieurs mots en majuscules.

**Fonctionnement :** Repère `(up)`, récupère le mot précédent et convertit toutes ses lettres en majuscules. Avec `(up, 3)`, elle transforme les trois mots précédents.

**Exemples :**

hello world (up) → hello WORLD
hello world (up, 2) → HELLO WORLD

### LOW : Mettre en minuscules

**Rôle :** Transforme un mot ou plusieurs mots en minuscules.

**Fonctionnement :** Repère `(low)`, récupère le ou les mots précédents, puis convertit leurs lettres en minuscules.

**Exemples :**

HELLO WORLD (low, 2) → hello world
HELLO WORLD (low) → HELLO world

### CAP : Mettre une majuscule initiale

**Rôle :** Transforme un mot pour que sa première lettre soit en majuscule et les suivantes en minuscules.

**Fonctionnement :** Repère `(cap)` et transforme le mot précédent. Avec `(cap, 3)`, elle applique la transformation aux trois mots précédents.

**Exemple :**

the golden gate bridge (cap, 3) → the Golden Gate Bridge

## 2. Les conversions numériques

### HEX : Convertir de l'hexadécimal en décimal

**Rôle :** Transforme un nombre écrit en base 16 en nombre décimal, c'est-à-dire en base 10.

**Fonctionnement :** Repère `(hex)`, récupère le nombre précédent, le convertit en base 10, puis remplace le nombre et le marqueur.

**Exemple :**

2A (hex) → 42

### BIN : Convertir du binaire en décimal

**Rôle :** Transforme un nombre écrit en base 2 en nombre décimal.

**Fonctionnement :** Repère `(bin)`, récupère le nombre binaire précédent, le convertit en base 10 et remplace les deux éléments.

**Exemple :**

1010 (bin) → 10

## 3. La ponctuation et les apostrophes

### PONCTUATION : Corriger la ponctuation

**Rôle :** Supprimer les espaces incorrects autour des signes de ponctuation.

**Fonctionnement :** Parcourt le texte pour coller les signes `.`, `,`, `!`, `?`, `:` et `;` au mot précédent et conserver un espace après le signe, sauf lorsque plusieurs signes forment un groupe.

**Exemple :**

hello ,world ! → hello, world!

**Attention :** `...`, `!?` ou `!!` doivent rester groupés, sans espace entre leurs signes.

### APOSTROPHES : Gérer les citations entre `' '`

**Rôle :** Retirer les espaces inutiles à l'intérieur des apostrophes qui encadrent une expression.

**Fonctionnement :** Repère une apostrophe ouvrante et sa correspondante fermante, puis supprime les espaces situés juste après la première et juste avant la seconde.

**Exemple :**

' hello world ' → 'hello world'

**Attention :** Dans le sujet, les apostrophes concernées sont des délimiteurs de citation. Il ne faut pas supprimer les apostrophes internes aux mots anglais, comme dans `isn't`.

## 4. La règle A → AN

### A-AN : Corriger l'article anglais

**Rôle :** Remplacer `a` par `an` lorsque le mot suivant commence par une voyelle (`a`, `e`, `i`, `o`, `u`) ou par `h`, conformément à la règle spécifique du sujet.

**Fonctionnement :**

1. Repérer le mot `a`.
2. Examiner le mot suivant.
3. Si sa première lettre correspond à la règle, remplacer `a` par `an`.
4. Sinon, conserver `a`.

**Exemples :**

a apple → an apple
a book → a book
a honest person → an honest person

## 5. TEST : Vérifier ton programme

**Rôle :** vérifier que les transformations fonctionnent correctement et que les règles ne se cassent pas entre elles.

**Fonctionnement :** Tu prépares un fichier d'entrée contenant des cas simples, puis des cas plus complexes. Tu lances ton programme et compares le fichier de sortie avec le résultat attendu.

Tu dois notamment tester :

- Chaque règle séparément.
- Les variantes avec un nombre, comme `(up, 3)`.
- Plusieurs règles dans la même phrase.
- Les espaces et la ponctuation.
- Les fichiers d'entrée et de sortie avec des noms différents.
- Les cas limites, comme un marqueur sans mot précédent.

Les tests ne transforment pas le texte : ils servent à détecter les erreurs.