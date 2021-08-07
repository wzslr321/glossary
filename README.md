<h3 align="center"> Personal CLI Glossary </h3>

#

I encounter new English words every single day, I decided to put bigger pressure 
on learning them. With <b> Glossary </b> you can add a new English word and its 
definition to the database, simply typing `aw my_word` in the terminal. 

#

---

<p align="center">
  Found it useful? Want more updates?
</p>

<p align = "center">
  <b> <i> Show your support by giving a :star: </b> </i>
</p>

---

<br>

<h3> How does it work? </h3>

Repository contains already built, executable file - `glossary`. 
It is the most convenient to create a new alias in your shell config file, e.g. `.zshrc`, 
`.bashrc`, `.bash_profile`. <br>
I added it to my .zshrc config with the following line: `alias aw="./development/words/glossary"`<br>
Depending on where did you clone this repository on your local machine, path my differ - change it as needed.

Now I can add new word (<i> hello </i> in this case) to the `new_words.txt` file, simply typing `aw hello` <br>
It runs without a docker container, so it is not stored in the database yet, and definition isn't assigned.

I can also run docker containers with the following command: `docker compose up` <br>
^ It requires <i> docker </i> and <i> docker compose </i> configured. 

Now with docker containers running, I can send a POST request to `localhost:8080/fetch`. <br>
It takes all words from new_words.txt and puts them into words.txt file.  <br>
It also scraps the web (https://www.dictionary.com) and saves its definition in the key-value database <br>
Of course, it also validates if words are being repeated and don't save them if so.

Send a GET request to `localhost:8080/<word>` to see definition of particular words. 
Send a POST request to `localhost:8080/<word>` to add a word and its definition to database. (Only with docker running)


<h3> Plans </h3>

Someday I wish to nicely display all the words on a website, currently I am not willing to 
code front-end though. Actual functionality is enough for my current needs.

<b> If someone is willing to code front-end, I can work on this project a little bit more 
and create much better API. </b>

Of course, <i> words.txt </i> are going to steady be updated with new words. 

---

<i> I am aware that code in this project is awful. I didn't want to spend too much time on 
building it. I just wanted to create a simple program to save a new words that I encounter 
while reading English books. It took me just few hours to create, without caring about 
architecture and clean code. </i>
