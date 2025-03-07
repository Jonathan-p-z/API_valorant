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
     - Cadence de tir
     - Capacité du chargeur
     - Temps de rechargement

4. **Liste des cartes** :
   - Affiche toutes les cartes disponibles dans le jeu.
   - Informations détaillées sur chaque carte :
     - Nom
     - Image de la carte

5. **Recherche et filtres** :
   - **Recherche** : Trouvez rapidement un agent ou une arme en entrant un mot-clé.
   - **Filtres** : Affinez les résultats selon des critères spécifiques (ex : rôle pour les agents, type d'arme).

6. **Affichage des résultats de recherche** :
   - Résultats dynamiques basés sur les mots-clés ou filtres appliqués.

7. **Page À propos** :
   - Présente des informations sur le projet et l'équipe de développement.

8. **Favoris** :
   - Ajoutez des agents, armes ou cartes à vos favoris pour un accès rapide.

---

## **Technologies Utilisées**
- **Langages** :
  - Go (Golang) : Back-end et logique métier.
  - HTML, CSS : Front-end et interface utilisateur.
  - JavaScript : Animation Front-end et interactivité.
  - JSON : Back-end pour le stockage des données.

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
| `/characters/details`   | Détails d'un agents spécifiques.       |
| `/characters/search`    | Recherche d'agents par mot-clé.        |
| `/weapons`              | Liste des armes Valorant.              |
| `/filtered_weapons`     | Filtrage des armes par critère.        |
| `/maps`                 | Liste des cartes Valorant.             |
| `/maps/details`         | Détails d'une carte spécifique.        |
| `/fav`                  | Liste des favoris de l'utilisateur.    |
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

### **Filtrage des armes :**
1. Naviguez vers `/weapons`.
2. Appliquez un filtre (ex : "Type d'arme = Fusils", "Cadence de tir = Au-dessus de 8").
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