import React from 'react'
import { NavLink } from "react-router-dom";
import Burger from "./SideBar";


const Nav: React.FC<{}> = () => {
  return (
    <>
            <nav className="navbar is-info hero-head">
                <div className="container">
                    <div className="navbar-brand">
                        <a className="navbar-item">
                            
                           <h1>Recepie App</h1> 
                        </a>
                        <span className="navbar-burger burger" data-target="navbarMenu">
                            <span></span>
                            <span></span>
                            <span></span>
                        </span>
                    </div>
                    <div id="navbarMenu" className="navbar-menu">
                        <div className="navbar-end">
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="/">
                                    <span className="icon">
                                        <i className="fa fa-home"></i>
                                    </span>
                                    <span>Home</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="keto">
                                    <span className="icon">
                                        <i className="fa fa-superpowers"></i>
                                    </span>
                                    <span>Keto</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="meat">
                                    <span className="icon">
                                        <i className="fa fa-book"></i>
                                    </span>
                                    <span>Meat</span>
                                </NavLink>
                            </span>
                            <span className="navbar-item">
                                <NavLink className="button is-white is-outlined" to="popular">
                                    <span className="icon">
                                        <i className="fa fa-github"></i>
                                    </span>
                                    <span>Popular</span>
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
                </div>
            </nav>
    </>
  )
}

export default Nav
