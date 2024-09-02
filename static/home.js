import { commentPage } from "./comment.js";
import {resetBody} from "./header.js";



async function fetchComments() {
    try {
        const response = await fetch('/recupPost');
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

const createPost = (pseudo, pp, text, datevalue, category, title, id ) => {
    const post = document.createElement('div');
    post.classList.add('post');


    const headerPost = document.createElement('div');
    headerPost.classList.add('headerPost');
    
    const contentPost = document.createElement('div');
    contentPost.classList.add('contentPost');

    const imgComment = document.createElement('img')
    imgComment.classList.add('imgComment')
    imgComment.src = "assets/comment.png"
    imgComment.addEventListener('click', function() {
        resetBody()
        commentPage(pseudo, pp, text, datevalue, category, id)
    })

    const profilePictureDiv = document.createElement('div')
    profilePictureDiv.classList.add('profilePictureDiv')
    
    const profilPicture = document.createElement('img');
    profilPicture.src = pp;
    profilPicture.classList.add('profilPicture');

    const user = document.createElement('p');
    user.classList.add('user');
    user.innerText = pseudo;

    const date = document.createElement('p');
    date.classList.add('date');
    date.innerText = datevalue;

    const categorie = document.createElement('p')
    categorie.classList.add('category')
    categorie.innerText = category

    const divTitle = document.createElement('div')
    divTitle.classList.add('divTitle')

    const titre = document.createElement('p')
    titre.classList.add('titlePost')
    titre.innerText = title

    const content = document.createElement('p');
    content.classList.add('content');
    content.innerText = text;

    post.appendChild(divTitle)
    divTitle.appendChild(titre)
    post.appendChild(headerPost);
    post.appendChild(contentPost);
    profilePictureDiv.append(profilPicture)
    headerPost.appendChild(profilePictureDiv);
    headerPost.appendChild(user);
    headerPost.appendChild(date);
    headerPost.appendChild(categorie);
    contentPost.appendChild(imgComment);
    contentPost.appendChild(content);
    return post;
};
export function homePage() {
    document.body.style.overflow = 'auto'
    
    const allPage = document.createElement('div');
    allPage.classList.add('allPage');
    
    const backgroundContainer = document.createElement('div');
    backgroundContainer.classList.add('backgroundContainer');
    
    const sectionPostContainer = document.createElement('div');
    sectionPostContainer.classList.add('sectionPostContainer');
    
    const postContainer = document.createElement('div');
    postContainer.classList.add('postContainer');
    const loadPosts = async () => {
        const comments = await fetchComments();
        if (comments){
            comments.forEach(element => {
                postContainer.append(createPost(element.username, element.pp, element.content, element.date, element.category, element.title, element.id));
            });
        }
    };
    const postSection = () => {
        const formDiv = document.createElement('form');
        formDiv.id = "sentMessage";
        
        const backPostSection = document.createElement('div');
        backPostSection.classList.add('backPostSection')
      
        const select = document.createElement('select');
        select.setAttribute('name', 'categories');
        select.setAttribute('id', 'categories');
        
        const placeHolderOption = document.createElement('option');
        placeHolderOption.text = 'Select a category';
        placeHolderOption.disabled = true;
        placeHolderOption.selected = true;

        select.appendChild(placeHolderOption)
        
        const options = [
            { value: 'Movie', text: 'Movie' },
            { value: 'Game', text: 'Game' },
            { value: 'Serie', text: 'Serie' },
            { value: 'Music', text: 'Music' },
            { value: 'Book', text: 'Book' },
            { value: 'Anime', text: 'Anime' }
        ];
        options.forEach(optionData => {
            const option = document.createElement('option');
            option.setAttribute('value', optionData.value);
            option.textContent = optionData.text;
            select.appendChild(option);
        });
        const writingTitle = document.createElement('input')
        writingTitle.classList.add('writingTitle')
        writingTitle.setAttribute('value', writingTitle.value)
        writingTitle.placeholder = "Add a title"
        writingTitle.required = true

        const writingPad = document.createElement('textarea');
        writingPad.classList.add('writingPad');
        writingPad.placeholder = "Write something you want to post...";
        writingPad.id = "textToSent";
        writingPad.required = true
        
        const buttonSectionPost = document.createElement('button');
        buttonSectionPost.classList.add('buttonSectionPost');
        buttonSectionPost.innerText = "Post";
        buttonSectionPost.type = "submit";
        formDiv.addEventListener('submit', async function(event) {
            event.preventDefault();
            const text = writingPad.value;
            var cat = select.value;
            if (cat == "Select a category") {
                cat = "No Category"
            }
            const title = writingTitle.value
            try {
                const response = await fetch('/createPost', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ text, cat, title })
                });
                const data = await response.json();
                if (data) {
                    writingPad.value = '';
                    writingTitle.value = '';
                    location.reload()
                } else {
                    console.error("Failed to post");
                }
            } catch (error) {
                console.error(error);
            }
        });
        backPostSection.append(select, writingTitle, writingPad, buttonSectionPost);
        formDiv.append(backPostSection);
        return formDiv;
    };
    sectionPostContainer.append(postSection());
    loadPosts();
    
    const bigLeftContainer = document.createElement('div');
    
    bigLeftContainer.classList.add('bigLeftContainer');
    const leftContainer = document.createElement('div');
    
    leftContainer.classList.add('leftContainer');
    const createCategorie = (text, id, imageUrl) => {
        const categorie = document.createElement('div');
        categorie.classList.add('categorie');
        categorie.innerText = text;
        categorie.id = id;
        const categorieImage = document.createElement('img');
        categorieImage.classList.add('categorieImage');
        categorieImage.src = imageUrl;
        categorie.alt = text;
        categorie.append(categorieImage);
        return categorie;
    };
    leftContainer.append(
        createCategorie("Movies", "categorie-1", "assets/movies.png"),
        createCategorie("Games", "categorie-2", "./assets/games.png"),
        createCategorie("Series", "categorie-3", "./assets/series.png"),
        createCategorie("Music", "categorie-4", "./assets/music.png"),
        createCategorie("Books", "categorie-5", "./assets/books.png"),
        createCategorie("Animes","categorie-6", "./assets/anime.png")
    );
    const scrollUp = document.createElement('div');
    scrollUp.id = 'scrollUp';
    const top = document.createElement('a');
    top.innerText = "Post";
    top.href = "#top";
    top.classList.add('toTop');
    jQuery(function() {
        $(function() {
            $(window).scroll(function() {
                if ($(this).scrollTop() > 320) {
                    $('#scrollUp').css('display', 'block');
                } else {
                    $('#scrollUp').removeAttr('style');
                }
            });
        });
    });
    scrollUp.append(top);
    bigLeftContainer.append(leftContainer, scrollUp);
    backgroundContainer.append(sectionPostContainer, postContainer);
    allPage.append(bigLeftContainer, backgroundContainer);
    document.body.append(allPage);
}