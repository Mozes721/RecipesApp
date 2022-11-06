import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import {useTheme} from "styled-components";

interface Props {
    text: string | number;
}

const SearchedReturn: React.FC<Props> = (props)=> {
    const [meals, setMeals] = useState([] as any[]);
    const [nutrients, setNutrients] = useState([] as any[])

    useEffect(() => {

        if (isNaN(props.text as number)) {
            getMealAsType();
        } else {
            getMealAsCaloire();
        }

    }, [props.text]);

    const getMealAsType = async () => {
        const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&type=${props.text}&number=1`)
        const data = await api.json();
        console.log(data.meals);
        setMeals(data.meals);
        setNutrients(data.nutrients);
        console.log(data.nutrients);
    }
    const getMealAsCaloire = async () => {
        const api = await fetch(`https://api.spoonacular.com/mealplanner/generate?apiKey=${process.env.REACT_APP_API_URL}&timeFrame=day&targetCalories=${props.text}&number=1`)
        const data = await api.json();
        console.log(data.meals);
        setMeals(data.meals);
        setNutrients(data.nutrients);
        console.log(data.nutrients);
    }
        if(meals !== []) {
            return (
                <main>
                <section>
                    <nav className="level">
                    { Object.entries(nutrients).map((nutrient,amount) =>
                        <div className="level-item has-text-centered">
                            <div>
                                <p className="heading">{nutrient[0]}</p>
                                <p className="title">{nutrient[1]}</p>
                            </div>
                        </div>)
                    }
                    </nav>
                </section>
                        <section className="container is-variable is-4">
                            <Splide options={{
                                perPage:3,
                                arrows:false,
                            }}>
                                {meals.map((meal)=> {
                                    return(
                                        <SplideSlide>
                                            <div className="card is-shady column is-10">
                                                <div className="card-content ">
                                                    <div className="content">
                                                        <h5>{meal.title}</h5>
                                                        <p className="mt-2 text-gray-800 text-sm">Ready in {meal.readyInMinutes}</p>
                                                        <p className="mt-2 text-gray-800 text-sm">Amount of servings {meal.servings}</p>
                                                        <a href={meal.sourceUrl} target="_blank" rel="noreferrer">
                                                            <span className="button is-info modal-button column is-narrow" data-target="modal-image2">Link</span>
                                                        </a>

                                                            <span className="button is-success column is-narrow" data-target="modal-image2">Add to Yumms</span>
                                                    </div>
                                                </div>
                                            </div>

                                        </SplideSlide>
                                    );
                                })}
                            </Splide>
                        </section>
                </main>
            )
        }
    return null;
}


export default SearchedReturn;