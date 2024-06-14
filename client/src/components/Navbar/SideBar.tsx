import React from "react";
import { slide as Menu } from "react-burger-menu";
import { NavLink } from "react-router-dom";
import styled from 'styled-components';
import { removeUser } from "../../hooks/removeUser";

const StyledBurgerMenu = styled(Menu)`
  /* Position and sizing of burger button */
  .bm-burger-button {
    position: fixed;
    width: 36px;
    height: 30px;
    left: 36px;
    top: 36px;
  }

  /* Color/shape of burger icon bars */
  .bm-burger-bars {
    background: #373a47;
  }

  /* Position and sizing of clickable cross button */
  .bm-cross-button {
    height: 24px;
    width: 24px;
  }

  /* Color/shape of close button cross */
  .bm-cross {
    background: #bdc3c7;
  }

  /* General sidebar styles */
  .bm-menu {
    background: #373a47;
    padding: 2.5em 1.5em 0;
    font-size: 1.15em;
  }

  /* Morph shape necessary with bubble or elastic */
  .bm-morph-shape {
    fill: #373a47;
  }

  /* Wrapper for item list */
  .bm-item-list {
    color: #b8b7ad;
    padding: 0.8em;
  }

  /* Individual item */
  .bm-item {
    display: inline-block;
  }

  /* Styling of overlay */
  .bm-overlay {
    background: rgba(0, 0, 0, 0.3);
  }
`;

interface BurgerProps {
    isAuthenticated: boolean | undefined;
    toggleBurgerMenu: boolean;
    isExpired: boolean;
}

const Burger: React.FC<BurgerProps> = ({ isAuthenticated, toggleBurgerMenu, isExpired }) => {
    return toggleBurgerMenu ? (
        <StyledBurgerMenu>
          <NavLink className="subtitle is-6 has-text-white" to="/">
                About Me
            </NavLink>
            {isAuthenticated && !isExpired ?
              null :
              (<NavLink className="subtitle is-6 has-text-white" to="/login-register">
                  Login/Register
              </NavLink>)
              }

            <NavLink className="subtitle is-6 has-text-white" to="keto-paelo">
                Keto&Paelo
            </NavLink>

            <NavLink className="subtitle is-6 has-text-white" to="whole-popular">
                Whole&Popular
            </NavLink>

            <NavLink className="subtitle is-6 has-text-white" to="vegan-vegetarian">
                Vegan&Vegetarian
            </NavLink>
            <span>
        {isAuthenticated ? (
            <div className="rows">
                <NavLink className="subtitle is-6 has-text-white" to="yumms">
                    Yumms
                </NavLink>
                <NavLink className="column ml-6 mt-6 is-6 has-text-danger" to="/" onClick={(e) => removeUser(e)}>
                    <span>Exit</span>
                </NavLink>
            </div>
        ) : null}
      </span>
        </StyledBurgerMenu>
    ) : null;
};

export default Burger;