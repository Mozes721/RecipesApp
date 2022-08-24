import React from "react";
import { slide as Menu } from "react-burger-menu";
import { NavLink } from "react-router-dom";

const Burger: React.FC<{}> = () => {
  return (
    <Menu>
      <NavLink className="menu-item" to="/">
        Home
      </NavLink>

      <NavLink className="menu-item" to="keto">
        Keto
      </NavLink>

      <NavLink className="menu-item" to="meat">
        Meat
      </NavLink>

      <NavLink className="menu-item" to="popular">
        Popular
      </NavLink>

      <NavLink className="menu-item" to="yumms">
        Yumms
      </NavLink>
    </Menu>
  );
};

export default Burger