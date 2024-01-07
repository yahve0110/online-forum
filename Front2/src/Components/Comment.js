import {
  CreateButtonTypeInnerTextClass,
  CreateElAddClassAddTextAddhref,
} from "../Helpers/CreateElements.js";
import {RenderPostPage} from "../Pages/PostPage.js";

export function Comment(author, content, created_at,comment_id,postId) {
  const commentContainer = CreateElAddClassAddTextAddhref(
    "div",
    "comment-container"
  );

  const commentContent = CreateElAddClassAddTextAddhref(
    "div",
    "comment-content",
    content
  );
  const commentAuthor = CreateElAddClassAddTextAddhref(
    "div",
    "comment-author",
    `Author: ${author}`
  );
  const commentCreatedAt = CreateElAddClassAddTextAddhref(
    "div",
    "comment-created-at",
    created_at
  );

  commentContainer.id = comment_id

  commentContainer.append(commentContent, commentAuthor, commentCreatedAt);
  if (author === localStorage.getItem("UserNickname")) {
    const deleteLink = CreateButtonTypeInnerTextClass(
      "button",
      "",
      "Delete",
      "comment-delete-btn"
    );
    deleteLink.addEventListener("click", (e) => deleteComment(e,postId));
    commentContainer.append(deleteLink);
  }
  return commentContainer;
}


async function deleteComment(e,postId){
 const commentId = e.target.parentNode.id
  try {
    const response = await fetch(`http://localhost:8080/delete-comment/${commentId}`, {
      method: "DELETE",
      credentials: "include",
    });

    if (response.ok) {
      console.log("Comment deleted successfully");
      RenderPostPage(postId)
      // Add any additional logic you need after deleting the comment
    } else {
      throw new Error(`Failed to delete comment: ${response.statusText}`);
    }
  } catch (error) {
    console.error("Error deleting comment:", error);
  }
}