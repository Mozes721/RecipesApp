import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";

interface Props {
    text: string | number;
}

const SearchedReturn: React.FC<Props> = (props)=> {
//     const [searched, setSearched] = useState<string | number>("")
//
//     useEffect(() => {
//     getSearched();
//     },[]);
//
//     const getSearched = async () => {
//     }
// };
    return (
        <section className="card-content">
            <h1>{props.text}</h1>
            {/*<div className="content">*/}
            {/*    <h5>{recepie.title}</h5>*/}
            {/*    <div className="columns is-mobile  is-centered is-one-quarter">*/}
            {/*        <a href={recepie.sourceUrl} target="_blank" rel="noreferrer">*/}
            {/*            <span className="button is-info modal-button column is-narrow" data-target="modal-image2">Link</span>*/}
            {/*        </a>*/}
            {/*        <span className="button is-success column is-narrow" data-target="modal-image2">Add to Yumms</span>*/}
            {/*    </div>*/}
        </section>
    )
}


export default SearchedReturn;