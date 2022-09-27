import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";

const Searched: React.FC = () => {
    return (
        <div className="has-addons-fullwidth">
            <div className="control">
                <input className="input" type="text" placeholder="Find a recipe" />
            </div>
            <div className="control">
                <a className="button is-info">
                    Search
                </a>
            </div>
        </div>
    )
}
export default Searched;