import { CreateButtonTypeInnerTextClass, CreateElAddClassAddTextAddhref } from "../Helpers/CreateElements.js";

export function Post(id, nickname, title, content, created_at) {
  if(content.length > 200){
    content = content.substring(0, 400) + "..."
  }
  const homePageContainer = document.querySelector(".homepage-container");

  const postContainer = CreateElAddClassAddTextAddhref("div", "post-container");

  const postTitle = CreateElAddClassAddTextAddhref("div", "post-title", title);
  const postContent = CreateElAddClassAddTextAddhref(
    "div",
    "post-content",
    content
  );
  const postAuthor = CreateElAddClassAddTextAddhref(
    "div",
    "post-author",
   `Author: ${nickname}`
  );
  const postCreatedAt = CreateElAddClassAddTextAddhref(
    "div",
    "post-created-at",
    created_at
  );



  postContainer.append(postTitle, postContent, postAuthor, postCreatedAt);

  homePageContainer.append(postContainer);

  postContainer.dataset.postId = id;

  return postContainer;
}
