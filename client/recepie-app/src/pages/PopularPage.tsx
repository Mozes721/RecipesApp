import { useEffect, useState } from "react";
// const dotenv = require('dotenv');

export function Popular() {
  const [popular, setPopular] = useState([] as any[]);

  useEffect(() => {
    getPopular();
  }, []);

const getPopular = async () => {
  const api = await fetch(`https://api.spoonacular.com/recipes/random?apiKey=${process.env.REACT_APP_API_URL}&number=10`)
  const data = await api.json();
  console.log(data.recipes)
  
  setPopular(data.recipes);
};
  return (
    <div>
      <div className="box cta">
      <p className="has-text-centered">
        <span className="tag is-primary">TOP</span> 10 Popular
      </p>
    </div>
    <section className="container">
        <div className="columns features">
          {popular.map((recepie, i) => {
            return(
              <div className="column is-4">
                <div className="card is-shady">
                  <div className="card-image">
                    <figure className="image is-4by3">
                      <img src={recepie.image} alt={recepie.title} className="modal-button" data-target="modal-image2" />
                    </figure>
                  </div>
                <div className="card-content">
                  <div className="content">
                    <h4>{recepie.title}</h4>
                    {/* <p>{recepie.summary}</p> */}
                    <span className="button is-link modal-button" data-target="modal-image2">Image modal</span>
                  </div>
                </div>
            </div>
        </div>
            )
          }
          )}
        </div>
    </section>
    </div>
  )
}
