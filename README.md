## Q53

This is q53 - as in "the question to the answer of 53" in reference to douglas adams number 42 : )
It is the basis to deliver question to the answers we have. It serves as a multipurpose tool to any question we might have.


Copyright (c) 2013. Manuel Klarmann. (mklarmann@gmail.com).
All rights reserved.  See the LICENSE file for license.




------------
### Install


install go

	download appengine sdk

adjusted GOPATH:

	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin

and added path/to/appengine to $PATH

added app.yaml with the following content:

	application: q53-xxxx # your appengine id
	version: 1
	runtime: go
	api_version: go1
	
	handlers:
	- url: /.*
	  script: _go_app


------------
### Work


test locally:

	dev_appserver.py mklarmann/

upload:

	appcfg.py update mklarmann/

format text:

	gofmt -w mklarmann

update old library version:

	go tool fix mklarmann

clear the datastore:

	dev_appserver.py --clear_datastore mklarmann/

------------
### Basic


Begin: challenge

Iterate through:

	action <-> reaction
	e.g.:
	question (Q) <-> answer (A)

End: satisfaction


Hence we initialize the program with the following set of answer and question:

	"Are you seeking for answers?" - ["yes","no"]
	"Are you satisfied?" - ["yes","no"]


------------
### Theory

We combine the principles of **information theory** (entropy), naive **Bayesian estimators** and the power of the **internet** for this tool.
We look at the likelihoods of a positive outcome to the questions as a distribution, where we want to decrease entropy (maximize information). Hence we choose the next question that will exactly do so, by being answerd. After the answer we calculate with our Bayesian estimator the adjusted probabilities. 

http://en.wikipedia.org/wiki/Conditional_entropy
http://en.wikipedia.org/wiki/Mutual_information
http://en.wikipedia.org/wiki/Bayesian_network


------------
### Example


Think of a person - guess your family version

	Question 1: The gender of the person is male?
	Question 2: The person is 20 years older than you?

	Question 3: Is it your Dad?
	Question 4: Is it your Mum?
	Question 5: Is it your brother?
	Question 6: Is it your sister?

Each question can be answered by ["yes","maybe","no"] the answer is A1 to A6 correspondingly.

Each question is valued by the answer to another question by probabilites

If the answer to Q3 is yes, and the person is sastified then:

Q1 has the following values for

	yes = 0.9	maybe =	0.05	no = 0.05

Q2 has the following values for

	yes = 0.8	maybe =	0.05	no = 0.15


Each answer is further filtered through the truth-matrix:

	Answer\Reality	|	yes		maybe	no
	--------------------------------------
	yes				|	0.8		0.1		0.1	
	maybe			|	0.1		0.75	0.15
	no				|	0.05	0.1		0.85




------------
### Evaluation


-> Find the best question Q_alpha:

	Q_alpha = max_i(sum_Q(Information_{qi}(Q(A)))

	I_{q1}(Q1) = I_{q1}(E_{q1})
	I_{q1}(Q3) = I(Q2 and Q3) + I_{q3}(E_{q3})

-> Receive the answer A_alpha:

			By A_alpha Q_alpha becomes the least interesting question.
			I_{q_alpha}(A_alpha) = I_{q_alpha}(Q_alpha)
			Truth-Of(A_alpha) ist not A_alpha


For each presented question there is an probability estimate p > 0.5 of the question being answered positive.
(We calculated the chance for X being the answer to your question by p, is that true?)


------------
### Training


Negative learning case:
The next selected question has a LOWER p than on the one before AND the answer before was NEGATIVE
-> we are missing a question, hence we ask:
"Can you tell me something special I might have missed about X?"

e.g.: 
The gender of the figure is male -> Is the gender of the figure mail?
He is Mr. X -> Is he Mr. X?

Yet how do we challenge a good question with a better one?

Each question is initialized with 1 vote for 0 knowledge of the other questions.
That is, there is the weight of 1 for the case of 0.5 tranition probabilities in case there are 3 possible answers. Each answer has the probability 0 and no vote.
With the first positive vote, this changes to accounting the transition probabilies half (as there is now one vote for transition and one for answer), and adding this 1 vote to one of the answers.

To not learn only on the question with the most wisdom and comprehensiveness already, the algorithm has to choose between best question from the information maximization sense and maximizing his wisdom (best learning rate). This ration is still to be found. (TODO)


Positive learning case:
The next selected question has a LOWER p than on the one before AND the answer before was POSITIVE
we ask the final question: "Are you satisfied?" - if the reply is "yes" - we start learning by adding votes to our path to either answer if it was predicted correctly, or the transitions, if they had to be used (because his answer had a human error)


------------
### Use Cases


1. Guess the person I am thinking of
2. What job should I do / who should I serve / where could I create the most value?
3. How sustainable is my food?
4. Who ist the person who could solve my problem?
5. ...


------------
### Cost of "Escaping the Maybe"


The relative time to response a question with a answer not maybe is the cost of escaping the maybe.
It can be used to assess the next question to be asked also by the cost it takes to answer them. So in case one wants a better confidence on answer delivered, we can point him the  cheapest question answered with a maybe to do research on. 


------------
### Modules


- Language: First question is what language should I address you?
- Maps: Where do you life, place a pin
- Databases: Market of jobs, people, wikipedia, co2 values
- Display: Display information alongside the next question (hence have more elaborate questions)
- Transfer: Allow for p2p distribution of the knowledge database

