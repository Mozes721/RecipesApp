import React, {useState, useEffect} from 'react';
import Searched from '../components/searched/Searched';
import SearchedResult from '../components/searched/SearchedResult';
import TableTypes from "../components/TableTypes";
const HomePage: React.FC = () => {
    const [text, setText] = useState<string | number>("");
    const [isModalOpen, setModalState] = React.useState(false);
    const toggleModal = () => setModalState(!isModalOpen);
    const callBack = (val: string | number) => {
        setText(val);
    }
  return (
      <section className="section py-4">
        <div className="container">
          <div className="column is-multiline is-centered">
            <div className="column has-text-centered">
              <span className="has-text-gray-dark">Welcome to My Recepie's</span>
              <h2 className="mt-2 mb-4 is-size-1 is-size-3-mobile has-text-weight-bold">Find  recipes of your liking based on Calories.</h2>
                <p className="subtitle has-text-gray mb-5">Otherwise you can click the Table <button className="button is-success" onClick={toggleModal}>here</button> and copy one of the meal options</p>
              <Searched text={callBack} />
            </div>
          </div>
            <TableTypes isOpen={isModalOpen} onClose={toggleModal} />
            <SearchedResult  text={text} />
        </div>
      </section>
  )
}

export default HomePage;