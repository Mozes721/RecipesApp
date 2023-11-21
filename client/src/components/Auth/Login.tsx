import React, { useState } from 'react';
import { AiOutlineMail } from "react-icons/ai";
import { RiLockPasswordFill, RiLoginBoxLine } from "react-icons/ri";
import { FcGoogle } from "react-icons/fc";
import { auth } from "../../config/firebase-config"
import { login } from "../../api/Auth/login";
import { signInWithGoogle } from "../../api/Auth/googleLogin";
import { useNavigate } from 'react-router-dom';

interface LoginFormProps {
  onSwitchForm: () => void;
}

const LoginForm: React.FC<LoginFormProps> = ({ onSwitchForm }) => {

  const [loginEmail, setLoginEmail] = useState("");
  const [loginPassword, setLoginPassword] = useState("");
  const navigate = useNavigate();

    const customLogin = (e: React.ChangeEvent<any>) => {
        login(auth, loginEmail, loginPassword)
            .then(() => {
                navigate("/yumms");
            })
            .catch((error) => {
                console.log(error.code, error.message);
            });
    }

    const googleSignIn = () => {
        signInWithGoogle(auth) // Pass your auth instance here
            .then(() => {
                navigate("/yumms");
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
            <input 
                className="input" 
                type="email" 
                placeholder="Email" 
                onChange={(event) => {
                      setLoginEmail(event.target.value);
                  }}
            />
            <AiOutlineMail className="icon is-small is-left" />
          </p>
        </div>
        <div className="field">
          <label className="label has-text-left">Password</label>
          <p className="control has-icons-left">
            <input 
                className="input" 
                type="password" 
                placeholder="Password" 
                onChange={(event) => {
                    setLoginPassword(event.target.value);
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
            <span onClick={(e) => customLogin(e)}>Login</span>
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
  
  
export default LoginForm;