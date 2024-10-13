import React, {useState} from 'react';
import Modal from './components/Modal';
import PTile from './components/PRT_Tile';

const App: React.FC = () => {
    const [isModalOpen, setModalOpen] = useState(false);

    const openModal = () => setModalOpen(true);
    const closeModal = () => setModalOpen(false);

    return (
        <div>
            <PTile label="Test Tile" onClick={openModal}/>
            <Modal isOpen={isModalOpen} onClose={closeModal}>
                <h2>Modal Title</h2>
                <p>This is the modal content.</p>
            </Modal>
        </div>
    );
};

export default App;