import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import {useTheme} from "styled-components";

interface Props {
    text: string | number;
}

const SearchedReturn: React.FC<Props> = (props)=> {
    const [searched, setSearched] = useState([] as any[]);
    useEffect(() => {
        getMeal();
    },[]);

    const getMeal = async () => {
        if (typeof props.text == "string") {
            const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&type=${props.text}&number=1`)
            const data = await api.json();
            setSearched(data.recepies);
        } else {
            const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&targetCalories=${props.text}&number=1`)
            const data = await api.json();
            setSearched(data.recepies);
        }

    }

    return (
        <div className="container">
            {searched.map(recepie) => {
                return (
                    <div>
            <section className="nutrients">
                <div className="level-item has-text-centered">
                    <div>
                        <p className="heading">Calories</p>
                        <p className="title">{searched.nutrients.calories.toFixed(0)}</p>
                    </div>
                </div>
                <div className="level-item has-text-centered">
                    <div>
                        <p className="heading">Following</p>
                        <p className="title">123</p>
                    </div>
                </div>
                <div className="level-item has-text-centered">
                    <div>
                        <p className="heading">Followers</p>
                        <p className="title">456K</p>
                    </div>
                </div>
                <div className="level-item has-text-centered">
                    <div>
                        <p className="heading">Likes</p>
                        <p className="title">789</p>
                    </div>
                </div>
            </section>

            <section>
            <div className="card is-shady column is-10">
                <div className="card-image">
                    <figure className="image is-3by2">
                        <img src={recepie.image} alt={recepie.title} className="modal-button" data-target="modal-image2" />
                    </figure>
                </div>
                <div className="card-content ">
                    <div className="content">
                        <h5>{recepie.title}</h5>
                        <div className="columns is-mobile  is-centered is-one-quarter">
                            <a href={recepie.sourceUrl} target="_blank" rel="noreferrer">
                                <span className="button is-info modal-button column is-narrow" data-target="modal-image2">Link</span>
                            </a>
                            <span className="button is-success column is-narrow" data-target="modal-image2">Add to Yumms</span>
                        </div>
                    </div>
                </div>
            </div>
        </section>
                </div>
            })}
        </div>

    )
}


export default SearchedReturn;