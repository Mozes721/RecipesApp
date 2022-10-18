import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import {useTheme} from "styled-components";

interface Props {
    text: string | number;
}

const SearchedReturn: React.FC<Props> = (props)=> {
    const [searched, setSearched] = useState("");

    console.log(searched);

    useEffect(() => {

        if (isNaN(props.text as number)) {
            getMealAsType();
        } else {
            getMealAsCaloire();
        }

    }, [props.text]);

    const getMealAsType = async () => {
        console.log(props.text);
        const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&type=${props.text}&number=1`)
        console.log(api)
        const data = await api.json();
        setSearched(data);
    }
    const getMealAsCaloire = async () => {
        const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&targetCalories=${props.text}&number=1`)
        console.log(api)
        const data = await api.json();
        setSearched(data);

    }
        if(searched !== "") {
            return (
                <main>
                <section className="nutritions">
                    <nav className="level">
                        <div className="level-item has-text-centered">
                            <div>
                                <p className="heading">Calories</p>
                                <p className="title"></p>
                            </div>
                        </div>
                        <div className="level-item has-text-centered">
                            <div>
                                <p className="heading">Protein</p>
                                <p className="title">123</p>
                            </div>
                        </div>
                        <div className="level-item has-text-centered">
                            <div>
                                <p className="heading">Fat</p>
                                <p className="title">456K</p>
                            </div>
                        </div>
                        <div className="level-item has-text-centered">
                            <div>
                                <p className="heading">Carbohydrates</p>
                                <p className="title">789</p>
                            </div>
                        </div>
                    </nav>
                </section>
                {/*<section className="recepie">*/}
                {/*    /!*{searched.map((recepie) => {*!/*/}
                {/*    /!*    return (*!/*/}
                {/*    /!*        <p>recepie</p>*!/*/}
                {/*    /!*    )*!/*/}
                {/*    /!*})}*!/*/}
                {/*</section>*/}
                </main>
            )
        }
        return null;
}


export default SearchedReturn;