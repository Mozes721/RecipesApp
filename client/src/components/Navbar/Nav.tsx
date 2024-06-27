import React, { useState, useEffect } from "react";
import { NavLink } from "react-router-dom";
import { AuthenticationUserStates } from '../../types/global'
import Burger from "./SideBar";
import {LuVegan} from "react-icons/lu";
import {useSelector} from "react-redux";
import { removeUser } from "../../hooks/removeUser";
import { House, FileText, Door, Book, Egg } from "phosphor-react";

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
                    {windowWidth > 1022 && (
                        <>
                            <div className="navbar-brand is-clickable">
                                <a
                                    onClick={(e) => setIsActive(!isActive)}
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
                        </>
                    )}

                </div>
                <div id="navbarMenu" className="navbar-menu">
                    <div className="navbar-end">
                            <span className="navbar-item">
                            <NavLink className="button is-white is-outlined" to="/">
                                    <FileText size={24} />
                                    <span>About Me</span>
                                </NavLink>
                            </span>
                        <span className="navbar-item">
                                {isAuthenticated ? null :
                                    <NavLink className="button is-white is-outlined" to="/login-register">
                                        <Door size={24} />
                                        <span>Login/Register</span>
                                    </NavLink>
                                }
                            </span>
                        <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="keto-paelo">
                                    <Egg size={24} />
                                    <span>Keto&Paelo</span>
                                </NavLink>
                            </span>
                        <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="whole-popular">
                                    <Book size={24} />
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
                                        <House size={24} />
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