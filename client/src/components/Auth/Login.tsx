import React, { useState } from 'react';
import { AiOutlineMail } from "react-icons/ai";
import { RiLockPasswordFill, RiLoginBoxLine } from "react-icons/ri";
import { FcGoogle } from "react-icons/fc";

interface LoginFormProps {
  onSwitchForm: () => void;
}


const LoginForm: React.FC<LoginFormProps> = ({ onSwitchForm }) => {

  const [loginEmail, setLoginEmail] = useState("");
  const [loginPassword, setLoginPassword] = useState("");

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
            <span>Login</span>
          </a>
          <a className="button is-link">
            <span className="icon is-small">
              <FcGoogle />
            </span>
            <span>Login with Google</span>
          </a>
        </div>
      </form>
    );
  };
  
  
export default LoginForm;