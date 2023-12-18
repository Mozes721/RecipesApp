import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import { AuthenticationUserStates, Recepie } from '../../types/global'
import { useSelector } from "react-redux";
import {addNewRecepie} from "../../api/Yumms/RecepieAdd";

const Whole: React.FC = () => {
    const [whole, setWhole] = useState([] as any[]);

    useEffect(() => {
        getWhole();
    }, []);

    const getWhole = async () => {

        const check = localStorage.getItem('whole');

        if(check) {
            setWhole(JSON.parse(check));
        } else {
            const api = await fetch(`https://api.spoonacular.com/recipes/random?apiKey=${process.env.REACT_APP_API_URL}&number=9&tags=whole30`)
            const data = await api.json();

            setWhole(data.recipes);

            localStorage.setItem('whole', JSON.stringify(data.recipes));
        }
    };

    const userID = useSelector((state: AuthenticationUserStates) => state.userID);
    const authToken = useSelector((state: AuthenticationUserStates) => state.authToken);

    const addRecepie = (title: string, url: string) =>  {
        const recepie: Recepie = {
            userID: userID,
            title: title,
            url: url,
            made: false,
            rating: 0
        };

        addNewRecepie(recepie, authToken);
    }

    return (
        <div>
            <div className="box cta">
                <p className="has-text-centered">
                    <span className="tag is-primary">Whole30 Recepies</span>
                </p>
            </div>
            <section className="container is-variable is-4">
                <h3>Whole Picks</h3>
                <Splide options={{
                    perPage:3,
                    arrows:false,
                }}>
                    {whole.map((recipe) => {
                        return(
                            <SplideSlide>
                                <div className="card is-shady column is-10">
                                    <div className="card-image">
                                        <figure className="image is-3by2">
                                            <img src={recipe.image} alt={recipe.title} className="modal-button" data-target="modal-image2" />
                                        </figure>
                                    </div>
                                    <div className="card-content">
                                        <div className="content">
                                            <h5>{recipe.title}</h5>
                                            <div className="columns buttons is-mobile  is-centered is-one-quarter">
                                                <a href={recipe.sourceUrl} target="_blank" rel="noreferrer">
                                                    <span className="button is-info modal-button column is-narrow" data-target="modal-image2">Link</span>
                                                </a>
                                                <span className="button is-success column is-narrow" data-target="modal-image2" onClick={() => addRecepie(recipe.title, recipe.sourceUrl)}>Add to Yumms</span>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                            </SplideSlide>
                        );
                    })}
                </Splide>
            </section>
        </div>
    )
}
export default Whole;