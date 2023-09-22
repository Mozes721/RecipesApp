import React, {useState} from 'react';
import { AiOutlineMail} from "react-icons/ai";
import { RiLockPasswordFill, RiLoginBoxLine} from "react-icons/ri";
import {FcGoogle} from "react-icons/fc"
const HomePage: React.FC = () => {
    const [text, setText] = useState<string | number>("");
    const [isModalOpen, setModalState] = React.useState(false);
    const toggleModal = () => setModalState(!isModalOpen);
    const callBack = (val: string | number) => {
        setText(val);
    }
  return (
    <section className="hero is-info is-fullheight has-background-dark">
    <div className="hero-body">
        <div className="container">
            <div className="columns is-centered">
                <div className="column is-three-fifths-desktop is-three-quarters-tablet has-text-centered">
                    <p className="bd-notification is-primary">
                        <p className="title">
                            Login
                        </p>
                        <div className="card">
                            <div className="card-content">
                                <div className="field">
                                  <label className="label has-text-left">Email</label>
                                    <p className="control has-icons-left has-icons-right">
                                        <input className="input" type="AiOutlineMail" placeholder="Email" />
                                        <AiOutlineMail className="icon is-small is-left"/>
                                    </p>
                                </div>
                                <div className="field">
                                  <label className="label has-text-left">Password</label>
                                    <p className="control has-icons-left">
                                        <input className="input" type="password" placeholder="Password" />
                                        <RiLockPasswordFill className="icon is-small is-left"/>
                                    </p>
                                </div>
                                <div className="buttons is-centered">
                                    <a className="button is-primary"><span className="icon is-small"><RiLoginBoxLine /></span><span>Login</span></a>
                                    <a className="button is-link"><span className="icon is-small"><FcGoogle /></span><span>Login with Google</span></a>
                                </div>
                                <hr />
                                <p><a href="#" className="has-text-link has-text-weight-bold">Not yet registered?</a></p>
                            </div>
                        </div>
                    </p>
                </div>
            </div>
        </div>
    </div>
</section>
  )
}

export default HomePage;