import { Comment } from "../Components/Comment.js";
import { CreateNavbar } from "../Components/Navbar.js";
import {
  CreateButtonTypeInnerTextClass,
  CreateElAddClassAddTextAddhref,
  CreateTextareaPlaceholderName,
} from "../Helpers/CreateElements.js";
import { Router } from "../index.js";





export const RenderPostPage = (postId) => {
  const container = document.querySelector(".root");
  container.innerHTML = "";

  // Get data if the user is logged in
  const isLogged = localStorage.getItem("isLogged");

  // Add the navbar
  CreateNavbar(isLogged);

  const separatePostPageContainer = CreateElAddClassAddTextAddhref(
    "div",
    "separatePostPageContainer"
  );
  const separatePostDataContainer = CreateElAddClassAddTextAddhref(
    "div",
    "separatePostDataContainer"
  );
  const separatePostDataCommentsContainer = CreateElAddClassAddTextAddhref(
    "div",
    "separatePostDataCommentsContainer"
  );
  const commentForm = CreateTextareaPlaceholderName(
    "add comment here",
    "comment-content"
  );
  commentForm.classList.add('comment-textarea')
  const addCommentBtn = CreateButtonTypeInnerTextClass(
    "button",
    "",
    "Add",
    "add-comment-btn"
  );

  const goBackBtn = CreateElAddClassAddTextAddhref(
    "a",
    "goBackBtn",
    "Go back",
    "/home"
  );


  addCommentBtn.addEventListener("click", (event) => {
    event.preventDefault(); // Prevent the default behavior
    addComment(postId, separatePostDataCommentsContainer);
    return false;
  });

  separatePostPageContainer.append(
    separatePostDataContainer,
    separatePostDataCommentsContainer,
    commentForm,
    addCommentBtn,
    goBackBtn
  );

  container.append(separatePostPageContainer);

  fetchPostData(postId)
    .then((postData) => {
      // Process the retrieved post data
      console.log("Post data:", postData);
      // Render the post on the page
      const postTitle = CreateElAddClassAddTextAddhref(
        "div",
        "post-title",
        `Title: ${postData.post.title}`
      );
      const postContent = CreateElAddClassAddTextAddhref(
        "div",
        "post-content",
        postData.post.content
      );
      const postAuthor = CreateElAddClassAddTextAddhref(
        "div",
        "post-author",
        postData.post.author_id
      );
      separatePostDataContainer.id = postData.post.id;
  
  
    if(postData.comments){
      postData.comments.forEach(comment => {
  
        separatePostDataCommentsContainer.append(Comment(comment.author, comment.content,comment.creation_time,comment.id))
        });
  
    }
      separatePostDataContainer.append(postTitle, postContent, postAuthor);
      if(postData.post.author_id === localStorage.getItem("UserNickname")){
        const deleteLink = CreateButtonTypeInnerTextClass("button","","Delete","post-delete-btn")
        deleteLink.addEventListener('click',(e)=>deletePost(e))
        separatePostDataContainer.append(deleteLink)
      }
    })
    .catch((error) => {
      // Handle errors, e.g., display an error message to the user
      console.error("Error handling post data:", error);
    });
};

async function fetchPostData(postId) {
  try {
    const response = await fetch(`http://localhost:8080/post/${postId}`, {
      method: "GET",
      credentials: "include",
    });

    if (response.ok) {
    
      const postData = await response.json();
  
      return postData;
    } else {
      throw new Error(`Failed to fetch post data: ${response}`);
    }
  } catch (error) {
    console.error("Error fetching post data:", error);
    throw error; // Propagate the error to the caller
  }
}

async function addComment(postId, separatePostDataCommentsContainer) {
  const form = document.querySelector(".separatePostPageContainer");
  const content = form.querySelector('textarea[name="comment-content"]').value;
  const commentData = { content: content };
  const textarea = document.querySelector('.comment-textarea');
  console.log(textarea);
  console.log(commentData);

  try {
    const response = await fetch(`http://localhost:8080/add-comment/${postId}`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(commentData),
    });

    if (response.ok) {
      const postData = await response.json();
      console.log(postData);
      if (postData.comments) {
        separatePostDataCommentsContainer.innerHTML = "";
     
        postData.comments.forEach(comment => {
          separatePostDataCommentsContainer.append(Comment(comment.author, comment.content, comment.creation_time,comment.id,postId));
          
        });
        // Clear the textarea
        textarea.value = '';

        document.body.scrollIntoView({ block: 'end'});

      }
    } else {
      throw new Error(`Failed to fetch post data: ${response.statusText}`);
    }
  } catch (error) {
    console.error("Error fetching post data:", error);
    throw error; // Propagate the error to the caller
  }
}






////////////////////


async function deletePost(e) {
  const postId = e.target.parentNode.id;

  try {
    const response = await fetch(`http://localhost:8080/delete-post/${postId}`, {
      method: "DELETE",
      credentials: "include",
    });
    window.location.hash = '#home'
    Router()
    if (response.ok) {
    

      console.log("Post deleted successfully");
  
    } else {
      throw new Error(`Failed to delete post: ${response.statusText}`);
    }
  } catch (error) {
    console.error("Error deleting post:", error);
  }
}
