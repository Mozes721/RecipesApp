import React, { useState, useEffect } from "react";
import { NavLink } from "react-router-dom";
import { AuthenticationUserStates } from '../../types/global'
import Burger from "./SideBar";
import {LuVegan} from "react-icons/lu";
import {useSelector} from "react-redux";
import { removeUser } from "../../hooks/removeUser";


const Nav: React.FC = () => {
  let [isActive, setIsActive] = React.useState<boolean>(false);
    const isAuthenticated = useSelector((state: AuthenticationUserStates) => state.authenticated);
    const email = useSelector((state: AuthenticationUserStates) => state.email);
    const [windowWidth, setWindowWidth] = useState<number>(window.innerWidth);

    useEffect(() => {
        const handleResize = () => {
            setWindowWidth(window.innerWidth);
        };

        window.addEventListener("resize", handleResize);

        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, []);

  return (
    <>
        <nav className="navbar is-info hero-head">
            <div className="container">
                {email ? (
                    <div className="navbar-item navbar-start">
                        <p className="has-text-white">Welcome, {email}</p>
                        {isAuthenticated && windowWidth > 1022 && (
                            <NavLink
                                className="button is-white is-outlined ml-2"
                                to="/login-register"
                                onClick={(e) => removeUser(e)}
                            >
                                <span>Exit</span>
                            </NavLink>
                        )}
                    </div>
                ) : (
                    <div className="navbar-item navbar-start">
                        <p className="has-text-white">Not Logged In</p>
                    </div>
                )}

                <div className="navbar-brand is-clickable">
                    <a
                        onClick={(e) =>  setIsActive(!isActive)}
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
                <Burger isAuthenticated={isAuthenticated} toggleBurgerMenu={isActive} />
            </div>
                    <div id="navbarMenu" className="navbar-menu">
                        <div className="navbar-end">
                            <span className="navbar-item">
                            <NavLink className="button is-white is-outlined" to="/">
                                    <span className="icon">
                                        <i className="fa-solid fa fa-file-text-o"></i>
                                    </span>
                                    <span>About Me</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="/login-register">
                                    <span className="icon">
                                        <i className="fa-solid fa-door-open"></i>
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
                                {isAuthenticated ? (
                                    <NavLink className="button is-white is-outlined" to="yumms">
                                        <span className="icon">
                                            <i className="fa fa-home"></i>
                                        </span>
                                        <span>Yumms</span>
                                    </NavLink>
                                    ) : null
                                }
                            </span>
                        </div>
                    </div>
            </nav>
    </>
  )
}
export default Nav
