import { resetBody } from "./header.js";
import { checkCommentPermission } from "./main.js"


var selectedchat = "general";
var conn = null;



if (checkCommentPermission) {
    connectToWebSocket("general");
}




async function fetchUser() {
    try {
        const response = await fetch('/recupUser');
        if (!response.ok) {
            throw new Error(`Network response was not ok: ${response.statusText}`);
        }
        const user = await response.json();
        return user;
    } catch (error) {
        console.error('Fetch operation failed:', error);
        return [];
    }
}


async function fetchMessage(name, id, limit) {
    try {
        let cookieValue = document.cookie
            .split('; ')
            .find(row => row.startsWith('sessionToken='))
            ?.split('=')[1];
        const response = await fetch('/loadMessage', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ cookie: cookieValue, username: name, thread: id.toString(), limit: limit.toString() })
        })
        if (!response.ok) {
            throw new Error(`Network response was not ok: ${response.statusText}`);
        }
        const message = await response.json();
        return message;
    } catch (error) {
        console.error('Fetch operation failed:', error);
        return [];
    }
}







function createChat(name, profilPicture, content, date) {
    console.log(profilPicture)
    const containers = document.getElementsByClassName('bigRightContainerChat');


    const container2 = containers[0];
    const children = container2.children;

    const container = children[1]


    if (container.length === 0) {
        console.error('Le conteneur avec la classe rightContainerChat est introuvable.');
        return;
    }





    if (isLastMessageFromUser("rightContainerChat", name)) {
        const newMessageContent = document.createElement('p');
        newMessageContent.classList.add('chatContent');
        newMessageContent.innerText = content;


        const chatElements = container.querySelectorAll('.chat');
        const lastChat = chatElements.length > 0 ? chatElements[chatElements.length - 1] : null;

        if (lastChat) {
            lastChat.appendChild(newMessageContent);
        }
        scrollToBottom(container);
        return lastChat;
    } else {
        const newChat = NewChat(name, profilPicture, content, date);
        container.appendChild(newChat);
        scrollToBottom(container);
        return newChat;
    }


}




function scrollToBottom(container) {
    container.scrollTop = container.scrollHeight 
    

}


function NewChat(name, profilPicture, content, date) {
    console.log("yo", profilPicture)
    const chatDate = document.createElement('p')
    chatDate.classList.add('chatDate')
    chatDate.innerText = date

    const chat = document.createElement('div');
    chat.classList.add('chat');

    const headerChat = document.createElement('div');
    headerChat.classList.add('headerChat');

    const chatPPDiv = document.createElement('div')
    chatPPDiv.classList.add('chatPPDiv')

    const chatPP = document.createElement('img');
    chatPP.classList.add('chatPP');
    chatPP.src = profilPicture

    const chatName = document.createElement('p');
    chatName.classList.add('chatName');
    chatName.innerText = name;

    const chatContent = document.createElement('p');
    chatContent.classList.add('chatContent');
    chatContent.innerText = content;

    chatPPDiv.append(chatPP)
    headerChat.append(chatPPDiv, chatName, chatDate);
    chat.append(headerChat, chatContent);
    
    return chat;
}



const writingChat = (name) => {
    const chatDiv = document.createElement('div')
    chatDiv.classList.add('chatDiv')

    const chatFormDiv = document.createElement('form')
    chatFormDiv.classList.add('chatFormDiv')

    const chatBar = document.createElement('input')
    chatBar.classList.add('chatBar')
    chatBar.placeholder = "Write a message"
    chatBar.addEventListener('input', () =>{
        sendMessage('istyping', 'istyping')
    })



    const chatButton = document.createElement('button')
    chatButton.classList.add('chatButton')
    chatButton.innerText = "Send"


    const typing = document.createElement('p')
    typing.classList.add('typing')
    typing.id = "typing"
    typing.innerText = `${name} is typing ...`
    typing.style.display = "none"

    chatDiv.append(typing,chatFormDiv )
    chatFormDiv.append(chatBar, chatButton)
    chatButton.addEventListener('click', function (event) {
        event.preventDefault()
        if (sendMessage(chatBar.value, "text")) {
            chatBar.value = ""

        } else {
            alert("error when you send a message")
        }
        


    })
    return chatDiv
}



export function chatPage() {


    selectedchat = "messagerie"
    changeChatRoom()

    document.body.style.overflow = 'hidden'

    const bigContainerChat = document.createElement('div')
    bigContainerChat.classList.add('bigContainerChat')

    const leftContainerChat = document.createElement('div')
    leftContainerChat.classList.add('leftContainerChat')

    const rightContainerChat = document.createElement('div')
    rightContainerChat.classList.add('rightContainerChat')




    const bigRightContainerChat = document.createElement('div')
    bigRightContainerChat.classList.add('bigRightContainerChat')



    const welcomeDiv = document.createElement('div')
    welcomeDiv.classList.add('welcomeDiv')
    welcomeDiv.id = 'welcomeDiv'
    welcomeDiv.innerHTML = `
        <div class="topWelcome">
            <h2 class="firstTitle"> There is nobody here ... </h2>
            <img class="monster1" src="assets/monster.png">
        </div>
        <div class="bottomWelcome">
            <h1 class="secondTitle"> Start a conversation with a friend ! </h1>
        </div>
        <img class="monster2" src="assets/monster2.png">
    `
    bigRightContainerChat.append(welcomeDiv)


    const createConv = (name, profilPicture, id) => {
        var lastMessage = false
        const friends = document.createElement('div')
        friends.classList.add('friends')
        friends.id = id

        const friendsPPDiv = document.createElement('div')
        friendsPPDiv.classList.add('friendsPPDiv')

        const friendsPP = document.createElement('img')
        friendsPP.classList.add('friendsPP')
        friendsPP.src = profilPicture

        const friendsOnline = document.createElement('div')
        friendsOnline.classList.add('friendsOnline')
        friendsOnline.style.backgroundColor = "red"
        friendsOnline.id = name

        const friendsName = document.createElement('p')
        friendsName.classList.add('friendsName')


        




        friendsName.innerText = name
        friendsPPDiv.append(friendsPP)
        friends.append(
            friendsPPDiv,
            friendsOnline,
            friendsName
        )


        friends.addEventListener("click", () => {
            var limit = 30
            bigRightContainerChat.innerHTML = ""
            rightContainerChat.innerHTML =""
            
            if (friends.id != "noThread") {
                selectedchat = friends.id
                changeChatRoom()
                bigRightContainerChat.append(writingChat(name))
                bigRightContainerChat.id = "rightchat" + friends.id

            } else {
                createChatFunction(name)


                    .then(ide => {
                        ide
                        friends.id = ide
                        selectedchat = ide
                        sendMessage(friends.id + "/" + friendsName.innerText, "click")
                        changeChatRoom()
                        bigRightContainerChat.id = "rightchat" + ide
                    });

                bigRightContainerChat.append(writingChat(name))
            }

            bigRightContainerChat.append(rightContainerChat)

    
            let debounceTimer;

            rightContainerChat.scrollTop = 1000

            rightContainerChat.addEventListener('scroll', () => {
                if (lastMessage || rightContainerChat.childElementCount == 0) {
                    return
                }
                clearTimeout(debounceTimer);
                debounceTimer = setTimeout(() => {
                    if (rightContainerChat.scrollTop <= 10) {
                        rightContainerChat.innerHTML = "";
                        limit += 30;

                        loadMessage(limit)
                            .then(el => {
                                rightContainerChat.scrollTop = 200;
                            });

                    }
                }, 200);

            });


            const loadMessage = async (limit = 20) => {

                const messages = await fetchMessage(name, friends.id, limit);
                if (messages) {
                    messages.forEach(element => {

                        if (element.lastmessage == "true") {
                            lastMessage = true
                        }

                        rightContainerChat.append(createChat(element.username, element.pp, element.message, element.date))
                        scrollToBottom(rightContainerChat);



                    });
                }
            };
            loadMessage(limit)


        })
        return friends
    }

    const loadPosts = async () => {
        const comments = await fetchUser();
        comments.forEach(element => {
            leftContainerChat.append(createConv(element.username, element.pp, element.thread))
        });
    };

    loadPosts()

    bigContainerChat.append(
        leftContainerChat,
        bigRightContainerChat
    )

    document.body.append(bigContainerChat)
}


async function createChatFunction(name) {
    try {
        let cookieValue = document.cookie
            .split('; ')
            .find(row => row.startsWith('sessionToken='))
            ?.split('=')[1];

        let response = await fetch('/createChat', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ cookie: cookieValue, username: name })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        let data = await response.json();
        return data.id;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
}


export function changeChatRoom(connected) {
    if (conn) {
        conn.close();
    }
    if (connected){
        connectToWebSocket(connected)
    } else {
        connectToWebSocket(selectedchat);
       
    }
}

function sendMessage(message, type) {
    var newMessage = message.trim();
    if (newMessage) {
        var formattedMessage;
        switch (type) {
            case 'text':
                formattedMessage = { type: 'text', content: newMessage };
                break;
            case 'ping':
                formattedMessage = { type: 'ping', content: newMessage };
                break;
            case 'click':
                formattedMessage = { type: 'click', content: newMessage };
                break;
            case 'istyping':
                formattedMessage = { type:'istyping', content: newMessage};
                break
            default:
                console.error('Type de message non reconnu',);
                return false;
        }
        try {
            var jsonMessage = JSON.stringify(formattedMessage);
            conn.send(jsonMessage);
            return true;
        } catch (error) {
            console.error('Erreur lors de la conversion en JSON ou de l\'envoi du message :', error);
            return false;
        }
    } else {
        return false;
    }
}


function connectToWebSocket(selectedchatId) {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws?room=" + selectedchatId)
    } else {
        alert("Your browser does not support WebSockets");
    }

    const receiveMessage = (msg) => {
        const chaty = document.getElementById("rightchat" + msg.thread)
        const chat = chaty.childNodes[1]
        if (chaty) {
            chat.append(createChat(msg.username, msg.pp, msg.message, msg.date))
            scrollToBottom(chat);
        }
        const chatToMove = document.querySelector(`div.friends[id="${msg.thread}"]`);
        const friendsDivs = document.querySelectorAll('.friends');
        if (chatToMove) {
            if (friendsDivs[0] !== chatToMove) {
                chatToMove.parentNode.removeChild(chatToMove);
                const firstDiv = friendsDivs[0];
                if (firstDiv) {
                    firstDiv.parentNode.insertBefore(chatToMove, firstDiv);
                } else {
                    document.querySelector('.friends').appendChild(chatToMove);
                }
            }
        } else {
            console.error(`Aucune div trouvée avec l'ID ${msg.thread}`);
        }


    }
    conn.onmessage = function (evt) {
        const message = JSON.parse(evt.data);
        switch (message.type) {
            case 'istyping':
                const typing = document.getElementById('typing')
                typing.style.display = "block"
                setTimeout(() => {
                    typing.style.display = "none"
                }, 3000)
                break
            case 'online':
                ChangeColorIfIsConnected(message.users)
                break;
            case 'message':
                receiveMessage(message);
                break;
            case 'roomChange':
                ChangeParentDivIdWithTheNameOfTheInnerText(message.name, message.idroom);
                break;
            case 'notif':
                switch (message.thread) {
                    case 'general':
                        const notifSignal = document.getElementById('notifSignal')
                        notifSignal.style.display = "block"
                        break
                    default:
                        Moove(message.thread)
                        break
                }

                break;
            default:
                break
        }
    };
}






function ChangeColorIfIsConnected(users) {
    const alluser = document.querySelectorAll('.friendsOnline')
    alluser.forEach(user => {
        if (user) {
            checkIf(user, users)
        }
    })
}






function checkIf(userDiv, users) {
    var Green = false
    users.forEach(userOnline => {
        if (userDiv.id == userOnline) {
            userDiv.style.backgroundColor = "green"
            Green = true
            return
        }
    })
    if (Green == false) {
        userDiv.style.backgroundColor = "red"
    }

}


function ChangeParentDivIdWithTheNameOfTheInnerText(name, newId) {
    const divs = document.querySelectorAll('.friendsName');
    divs.forEach(div => {
        if (div.innerText.trim() === name.trim()) {
            if (div.parentElement) {
                div.parentElement.id = newId;
            }
        }
    });
}


function Moove(threadid) {
    const chatToMove = document.querySelector(`div.friends[id="${threadid}"]`);
    const friendsDivs = document.querySelectorAll('.friends');
    if (chatToMove) {
        if (friendsDivs[0] !== chatToMove) {
            chatToMove.parentNode.removeChild(chatToMove);
            const firstDiv = friendsDivs[0];
            if (firstDiv) {
                firstDiv.parentNode.insertBefore(chatToMove, firstDiv);
            } else {
                document.querySelector('.friends').appendChild(chatToMove);
            }
        }
    } else {
        console.error(`Aucune div trouvée avec l'ID ${threadid}`);
    }
}

function isLastMessageFromUser(containerClass, userName) {
    const container = document.querySelector(`.${containerClass}`);

    if (!container) {
        console.error(`Aucun conteneur trouvé avec la classe ${containerClass}`);
        return false;
    }

    const chatElements = container.querySelectorAll('.chat');


    if (chatElements.length === 0) {
        return false;
    }


    const lastChat = chatElements[chatElements.length - 1];

    const lastUserName = lastChat.querySelector('.chatName').textContent.trim();

    return lastUserName === userName;
}
