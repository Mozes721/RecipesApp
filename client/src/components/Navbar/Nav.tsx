import React from "react";
import { NavLink } from "react-router-dom";
import AuthenticationUserStates from '../../types/global.d'
import Burger from "./SideBar";
import {LuVegan} from "react-icons/lu";
import {useSelector} from "react-redux";


const Nav: React.FC<{}> = () => {
  const [isActive, setisActive] = React.useState<boolean>(false);
    const email = useSelector((state: AuthenticationUserStates) => state.email);

  return (
    <>
            <nav className="navbar is-info hero-head">
                <div className="container">
                     <div className="navbar-brand ">
                        <a
                            onClick={() => {
                            setisActive(!isActive);
                            }}
                            role="button"
                            className={`navbar-burger burger`}
                            aria-label="menu"
                            aria-expanded="false"
                            data-target="navbarMenu"
                        >
                            <span aria-hidden="true"></span>
                            <span aria-hidden="true"></span>
                            <span aria-hidden="true"></span>
                        </a>
                        </div>
                        <Burger />
                        </div>
                    <div id="navbarMenu" className="navbar-menu">
                        <div className="navbar-start">
                            <span className="navbar-item">
                                {email ? (
                                    <p className="has-text-white">Welcome, {email}</p>
                                ) : (
                                    <p className="has-text-white">Not Logged In</p>
                                )}
                            </span>
                        </div>
                        <div className="navbar-end">
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="/">
                                    <span className="icon">
                                        <i className="fa-solid fa-home"></i>
                                    </span>
                                    <span>Login/Register</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="keto-paelo">
                                    <span className="icon">
                                        <i className="fa fa-superpowers"></i>
                                    </span>
                                    <span>Keto&Paelo</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="whole-popular">
                                    <span className="icon">
                                        <i className="fa fa-book"></i>
                                    </span>
                                    <span>Whole&Popular</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined p-3" to="vegan-vegetarian">
                                    <LuVegan />
                                    <span className="ml-1">Vegan&Vegetarian</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="yumms">
                                    <span className="icon">
                                        <i className="fa fa-home"></i>
                                    </span>
                                    <span>Yumms</span>
                                </NavLink>
                            </span>
                        </div>
                    </div>
            </nav>
    </>
  )
}
export default Nav
