import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import { AuthenticationUserStates } from '../../types/global'
import { useSelector } from "react-redux";
import { removeUser } from "../../hooks/removeUser";

const Keto: React.FC = () => {
  const [keto, setKeto] = useState([] as any[]);

  useEffect(() => {
    getKeto();
  }, []);

const getKeto = async () => {

  const check = localStorage.getItem('keto');
  
  if(check) {
    setKeto(JSON.parse(check));
  } else {
  const api = await fetch(`https://api.spoonacular.com/recipes/random?apiKey=${process.env.REACT_APP_API_URL}&number=9&tags=ketogenic`)
  const data = await api.json();

  console.log(data.recipes)
  setKeto(data.recipes);

  localStorage.setItem('keto', JSON.stringify(data.recipes));
  }
};

const addRecepie = (title: string, url: string) =>  {

}

  return (
    <div>
      <div className="box cta">
      <p className="has-text-centered">
        <span className="tag is-primary">Keto Recipes</span>
      </p>
    </div>
    <section className="container is-variable is-4">
        <h3>Popular Picks</h3>
          <Splide options={{
            perPage:3,
            arrows:false,
          }}>
          {keto.map((recepie) => {
            return(
              <SplideSlide>
                
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
                        <span className="button is-success column is-narrow" data-target="modal-image2" onClick={() => addRecepie(recepie.title, recepie.sourceUrl)}>Add to Yumms</span>
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
export default Keto;