import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";

interface Props {
    text: any;
}

const Searched: React.FC<Props> = (props) => {
    const [input, setInput] = useState<string | number>('');
    const handleSubmit = (event: any) => {
        props.text(input);
        event.preventDefault();
        event.target.reset();
}
    return (
        <form onSubmit={handleSubmit} >
            <div className="has-addons-fullwidth">
                <div className="control">
                    <input className="input" type="text" placeholder="Input calories like 2000 or recipes type from table" onChange={e => setInput(e.target.value)} required/>
                </div>
                    <div className="control 0.75rem">
                    <button className="button is-info">
                        Search
                    </button>
                </div>
            </div>
        </form>
    )
}
export default Searched;