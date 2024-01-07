export function CreateElAddClassAddTextAddhref(
  element,
  elClass = "",
  text = "",
  href = ""
) {
  const resultElement = document.createElement(element);
  if (elClass != "") {
    resultElement.classList.add(elClass);
  }
  if (text != "") {
    resultElement.innerText = text;
  }
  if (href != "") {
    resultElement.href = href;
  }
  return resultElement;
}

export function CreateInputTypePlaceholderName(
  elementType,
  inputType,
  placeholder,
  name
) {
  const inputElement = document.createElement(elementType);

  // Check if the element is a textarea
  if (elementType.toLowerCase() === "textarea") {
    // Use "textContent" property for textarea content
    inputElement.textContent = placeholder;
  } else {
    // Set type for other input elements
    inputElement.type = inputType;
    // Set placeholder and name attributes
    inputElement.placeholder = placeholder;
    inputElement.name = name;
  }

  return inputElement;
}

export function CreateButtonTypeInnerTextClass(
  el,
  type = "",
  innerText = "",
  elClass = "btn"
) {
  const resultElement = document.createElement(el);
  resultElement.type = type;
  resultElement.innerText = innerText;
  resultElement.classList.add(elClass);
  return resultElement;
}

export function CreateTextareaPlaceholderName(placeholder, name) {
  const textarea = document.createElement("textarea");
  textarea.placeholder = placeholder;
  textarea.name = name;
  return textarea;
}
