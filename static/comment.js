export function commentPage(pseudo, pp, text, datevalue, category, id) {
    const bigContainerComment = document.createElement('div')
    bigContainerComment.classList.add('bigContainerComment')

    const secondContainerComment = document.createElement('div')
    secondContainerComment.classList.add('secondContainerComment')

    const postComment = document.createElement('div')
    postComment.classList.add('postComment')
    
    const headerPostComment = document.createElement('div')
    headerPostComment.classList.add('headerPostComment')

    const profilePictureComment = document.createElement('div')
    profilePictureComment.classList.add('profilePictureComment')

    const profilePicture = document.createElement('img')
    profilePicture.classList.add('profilPicture')
    profilePicture.src = pp

    const name = document.createElement('p')
    name.classList.add('name')
    name.textContent = pseudo

    const date = document.createElement('p')
    date.classList.add('date')
    date.textContent = datevalue

    const cate = document.createElement('p')
    cate.classList.add('cate')
    cate.textContent = category

    const contentPostComment = document.createElement('div')
    contentPostComment.classList.add('contentPostComment')

    const content = document.createElement('p')
    content.classList.add('content')
    content.textContent = text

    const formDiv = document.createElement('form');
    formDiv.id = "sentComment";


    const commentArea = document.createElement('div')
    commentArea.classList.add('commentArea')

    const textCommentArea = document.createElement('textarea')
    textCommentArea.classList.add('textCommentArea')
    textCommentArea.placeholder = 'Post your reply'
    textCommentArea.required = true

    const buttonCommentArea = document.createElement('button')
    buttonCommentArea.classList.add('buttonCommentArea')
    buttonCommentArea.innerText = 'Post'


    const IdToSent = id
    formDiv.addEventListener('submit', async function(event) {
        event.preventDefault();
        const texte = textCommentArea.value;

        try {
            const response = await fetch('/sentComment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ id :IdToSent, text: texte, })
            });
            const data = await response.json();
            if (data) {
                resetBody()
                commentPage(pseudo, pp, text, datevalue, category, id)


            } else {
                console.error("Failed to post");
            }
        } catch (error) {
            console.error(error);
        }
    });

   


    const allComments = document.createElement('div')
    allComments.classList.add('allComments')


    commentArea.append(textCommentArea, buttonCommentArea)
    formDiv.append(commentArea)

    profilePictureComment.append(profilePicture)

    headerPostComment.append(profilePictureComment, name, date, cate)
    contentPostComment.append(content)
    postComment.append(
        headerPostComment, 
        contentPostComment
    )
    const loadPosts = async () => {
        const comments = await fetchComments(IdToSent);
        if (comments){
            comments.forEach(element => {
                console.log(element)
                allComments.append(everyComment(element.content, element.pp, element.username, element.date))
            });
        }
    };
    loadPosts()


    secondContainerComment.append(postComment, formDiv, allComments)
    bigContainerComment.append(secondContainerComment)
    document.body.append(bigContainerComment)
}







async function fetchComments(id) {
    try {
        const response = await fetch('/recupComments', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ text:"ignore", id:id})
        });
        if (!response.ok) {
            throw new Error(`Network response was not ok: ${response.statusText}`);
        }
        const comments = await response.json();
        return comments;
    } catch (error) {
        console.error('Fetch operation failed:', error);
        return [];
    }
}



const everyComment = (content, pp, pseudo, dateValue) => {
    const comment = document.createElement('div')
    comment.classList.add('comment')

    const headerComment = document.createElement('div')
    headerComment.classList.add('headerComment')

    const contentComment = document.createElement('div')
    contentComment.classList.add('contentComment')
    contentComment.innerText = content

    const ppCommentDiv = document.createElement('div')
    ppCommentDiv.classList.add('ppCommentDiv')

    const ppComment = document.createElement('img')
    ppComment.classList.add('ppComment')
    ppComment.src = pp

    const userComment = document.createElement('p')
    userComment.classList.add('userComment')
    userComment.innerText = pseudo

    const dateComment = document.createElement('p')
    dateComment.classList.add('dateComment')
    dateComment.innerText = dateValue

    ppCommentDiv.append(ppComment)
    headerComment.append(ppCommentDiv, userComment, dateComment)
    comment.append(headerComment, contentComment)
    return comment
}
export function resetBody(){
    const elementsToRemove = [...document.body.children].filter(el => el.tagName !== 'HEADER');
    elementsToRemove.forEach(el => el.remove());
}
