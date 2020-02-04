import { required, email, max, min } from "vee-validate/dist/rules";
import { extend } from "vee-validate";

extend("required", {
  ...required,
  message: "Данное поле должно быть заполненно"
});

extend("max", {
  ...max,
    message: "Данное поле может содержать не более {length} знаков"
});

extend("min", {
  ...min,
    message: "Данное поле должно содержать {length} знаков или более"
});

extend("email", {
  ...email,
  message: "Некорректный адрес электронной почты"
});

extend('password', {
  params: ['target'],
  validate(value, { target }) {
    return value === target;
  },
  message: 'Пароли не совпадают'
});