# **API Valorant Tracker**

## **Description**
`API Valorant Tracker` est un projet web conçu pour explorer l'univers de Valorant, le célèbre jeu de tir tactique développé par Riot Games. Ce projet propose une interface intuitive pour consulter des informations détaillées sur les agents, les armes et d'autres éléments du jeu. Toutes les données sont récupérées dynamiquement via l'API officielle de Valorant, garantissant des informations à jour et précises.

---

## **Fonctionnalités**
### **Pages Principales :**
1. **Page de connexion** :
   - Permet aux utilisateurs de se connecter pour accéder à leur expérience personnalisée.
   - Simple et sécurisée.
   
2. **Liste des agents (Personnages)** :
   - Affiche tous les agents disponibles dans le jeu.
   - Informations détaillées sur chaque agent :
     - Nom
     - Description
     - Image de l'agent
     - Rôle (ex : Duelist, Sentinel, etc.)

3. **Liste des armes** :
   - Détail des armes disponibles dans Valorant.
   - Informations affichées :
     - Nom
     - Type d’arme (Fusil, Pistolet, etc.)
     - Image de l’arme.

4. **Recherche et filtres** :
   - **Recherche** : Trouvez rapidement un agent ou une arme en entrant un mot-clé.
   - **Filtres** : Affinez les résultats selon des critères spécifiques (ex : rôle pour les agents).

5. **Affichage des résultats de recherche** :
   - Résultats dynamiques basés sur les mots-clés ou filtres appliqués.

6. **Page À propos** :
   - Présente des informations sur le projet et l'équipe de développement.

---

## **Technologies Utilisées**
- **Langages** :
  - Go (Golang) : Back-end et logique métier.
  - HTML, CSS : Front-end et interface utilisateur.
  - JS : Animation Front-end et animation .
  - Json : Back-end pour le stockage des données.

- **APIs** :
  - API officielle de Valorant pour récupérer les données dynamiques :
    - [Agents Endpoint](https://valorant-api.com/v1/agents)
    - [Weapons Endpoint](https://valorant-api.com/v1/weapons)
    - [Maps Endpoint](https://valorant-api.com/v1/maps)

- **Frameworks et outils** :
  - Go `net/http` : Gestion du serveur web.
  - Go `html/template` : Rendu des templates HTML.
  - Go `encoding/json` : Manipulation des données JSON.

---

## **Installation et Configuration**
### **Prérequis**
- **Go** : Assurez-vous que Go est installé sur votre système (version 1.20 ou supérieure).
- **Connexion Internet** : Nécessaire pour récupérer les données via l'API Valorant.

### **Étapes d'installation**
1. Clonez le projet :
   ```bash
   git clone https://github.com/Jonathan-p-z/api_valorant.git
   cd api_valorant
   ```

2. Initialisez les modules Go :
   ```bash
   go mod tidy
   ```

3. Lancez le serveur :
   ```bash
   go run main.go
   ```

4. Ouvrez votre navigateur et accédez à :
   ```
   http://localhost:8080/login
   ```

---

## **Fonctionnement**
### **Routes Principales**
| URL                     | Description                            |
|-------------------------|----------------------------------------|
| `/login`                | Page de connexion.                     |
| `/characters`           | Liste des agents Valorant.             |
| `/characters/search`    | Recherche d'agents par mot-clé.        |
| `/characters/filter`    | Filtrage des agents par critère.       |
| `/weapons`              | Liste des armes Valorant.              |
| `/maps`                 | Liste des cartes Valorant.             |
| `/maps/details`         | Détails d'une carte spécifique.        |
| `/about`                | Page À propos du projet.               |

---

## **Exemple de Fonctionnalités**

### **Recherche d'un agent :**
1. Naviguez vers `/characters/search`.
2. Entrez un mot-clé (ex : "Jett").
3. Obtenez les résultats correspondants.

### **Filtrage des agents :**
1. Naviguez vers `/characters/filter`.
2. Appliquez un filtre (ex : "Rôle = Duelist").
3. La liste sera mise à jour automatiquement.

---

## **Améliorations Futures**
- Ajouter une base de données pour la persistance des favoris utilisateur.
- Intégrer une authentification plus robuste avec JWT.
- Ajouter une page pour afficher les cartes (maps) du jeu.
- Rendre l'interface plus interactive avec JavaScript.

---

## **Contribution**
Les contributions sont les bienvenues ! Pour contribuer :
1. Forkez le dépôt.
2. Créez une branche pour vos modifications.
3. Soumettez une pull request.

---

## **Auteur**
Ce projet a été réalisé par **https://github.com/Jonathan-p-z** dans le cadre d'un projet pédagogique.