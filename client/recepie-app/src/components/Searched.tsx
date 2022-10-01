import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";

interface InputValue {
    text: any;
}

const Searched: React.FC<InputValue> = (props) => {
    const [input, setInput] = useState<string | number>('');
    const getInputValue = () => {
        props.text.value = input;
        console.log(props.text.value);

    }
    return (
        <div className="has-addons-fullwidth">
            <div className="control">
                <input className="input" type="text" placeholder="Find a recipe" onChange={e => setInput(e.target.value)}/>
            </div>
            <div className="control">
                <button className="button is-info" onClick={getInputValue}>
                    Search
                </button>
            </div>
        </div>
    )
}
export default Searched;