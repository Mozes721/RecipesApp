import React, { useState } from 'react';
import { AiOutlineMail} from "react-icons/ai";
import { RiLockPasswordFill, RiLoginBoxLine} from "react-icons/ri"

import {FcGoogle} from "react-icons/fc"
import { auth } from "../../config/firebase-config";
import { signIn } from "../../api/Auth/signin";
import {postUserID} from "../../api/Auth/postUserID";
import { useNavigate } from 'react-router-dom';
import {signInWithGoogle} from "../../api/Auth/googleLogin";

interface RegisterFormProps {
    onSwitchForm: () => void;
  }



  const RegisterForm: React.FC<RegisterFormProps> = ({ onSwitchForm }) => {

    const [registerEmail, setRegisterEmail] = useState("");
    const [registerPassword, setRegisterPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const navigate = useNavigate();

    const customRegister = (e: React.ChangeEvent<any>) => {
          e.preventDefault();
          if (registerPassword !== confirmPassword) {
              alert("Passwords don't match")
          } else {
              signIn(auth, registerEmail, registerPassword)
                  .then((user) => {
                      navigate("/yumms");
                      postUserID(user);
                  })
                  .catch((error) => {
                      const errorCode = error.code;
                      const errorMessage = error.message;
                      console.log(errorCode, errorMessage);
                      
                  });
          }
      }

      const googleSignIn = () => {
          signInWithGoogle(auth) // Pass your auth instance here
              .then((user) => {
                  // Successful sign-in, navigate to "/yumms"
                  navigate("/yumms");
                  console.log(user);
              })
              .catch((error) => {
                  console.log(error.code, error.message);
              });
      }

    return (
      <form>
        <div className="field">
          <label className="label has-text-left">Email</label>
          <p className="control has-icons-left has-icons-right">
            <input className="input" type="AiOutlineMail" placeholder="Email" 
                onChange={(event) => {
                    setRegisterEmail(event.target.value);
                }}
            />
            <AiOutlineMail className="icon is-small is-left" />
          </p>
        </div>
        <div className="field">
          <label className="label has-text-left">Password</label>
          <p className="control has-icons-left">
            <input className="input" type="password" placeholder="Password" 
                onChange={(event) => {
                    setRegisterPassword(event.target.value);
                }}
            />
            <RiLockPasswordFill className="icon is-small is-left" />
          </p>
        </div>
        <div className="field">
          <label className="label has-text-left">Repeat password</label>
          <p className="control has-icons-left">
            <input className="input" type="password" placeholder="Repeat Password" 
                 onChange={(event) => {
                    setConfirmPassword(event.target.value);
                }}
            />
            <RiLockPasswordFill className="icon is-small is-left" />
          </p>
        </div>
        <div className="buttons is-centered">
          <a className="button is-primary">
            <span className="icon is-small">
              <RiLoginBoxLine />
            </span>
            <span onClick={(e) => customRegister(e)}>Register</span>
          </a>
          <a className="button is-link">
            <span className="icon is-small">
              <FcGoogle />
            </span>
            <span onClick={googleSignIn}>Login with Google</span>
          </a>
        </div>
      </form>
    );
  };


export default RegisterForm;
