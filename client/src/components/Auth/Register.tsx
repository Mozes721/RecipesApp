import React, { useState } from 'react';
import { AiOutlineMail} from "react-icons/ai";
import { RiLockPasswordFill, RiLoginBoxLine} from "react-icons/ri"

import {FcGoogle} from "react-icons/fc"

interface RegisterFormProps {
    onSwitchForm: () => void;
  }



  const RegisterForm: React.FC<RegisterFormProps> = ({ onSwitchForm }) => {

    const [registerEmail, setRegisterEmail] = useState("");
    const [registerPassword, setRegisterPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");


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
            <span>Register</span>
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


export default RegisterForm;
