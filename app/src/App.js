import {
  AlertDialog,
  AlertDialogBody,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogOverlay,
  Box,
  Button,
  Card,
  CardBody,
  CardHeader,
  ChakraProvider,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Select,
  StackDivider,
  Text,
  VStack,
  theme,
  useDisclosure
} from '@chakra-ui/react';
import React, { useEffect, useRef, useState } from 'react';

import { ColorModeSwitcher } from './ColorModeSwitcher';

function CarFormModal({isEdit, car, colorOptions, setCars}) {
  const { isOpen, onOpen, onClose } = useDisclosure();  
  const formRef = useRef();

  async function addCar(e) {
    e.preventDefault();
    
    const response = await fetch('/api/car/create', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        make     : e.target.make.value,
        model    : e.target.model.value,
        color    : parseInt(e.target.color.value),
        buildDate: new Date(e.target.buildDate.value).toISOString()
      })
    });
    
    if (response.status === 200) {
      const result = await response.json();
      setCars(prev => [...prev, result.result]);

      onClose();
    }
  }  

  async function editCar(e) {
    e.preventDefault();
    
    const response = await fetch(`/api/car/${car.id}`, {
      method: 'PUT',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id       : parseInt(e.target.id.value),
        make     : e.target.make.value,
        model    : e.target.model.value,
        color    : parseInt(e.target.color.value),
        buildDate: new Date(e.target.buildDate.value).toISOString()
      })
    });

    
    if (response.status === 200) {
      window.location.reload();
    }
  }

  function getDefaultDate() {    
    return car?.buildDate ? new Date(car.buildDate).toISOString() : "";
  }    

  return (
    <>
      <Button colorScheme="blue" variant={isEdit ? "outline" : "solid"} onClick={onOpen}>
        {isEdit ? 'Edit car' : 'Add car'}
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Create Car</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <form onSubmit={isEdit ? editCar : addCar} ref={formRef}> 
              <VStack gap={3}>
                <Input type='hidden' name='id' value={car?.id} />
                <FormControl>
                  <FormLabel>Make</FormLabel>
                  <Input name='make' type='text' defaultValue={car?.make} />      
                </FormControl>
                <FormControl>
                  <FormLabel>Model</FormLabel>
                  <Input name='model' type='text' defaultValue={car?.model}  />      
                </FormControl>
                <FormControl>
                  <FormLabel>Color</FormLabel>
                  <Select name='color' placeholder='Select color' defaultValue={`${car?.colorId}`}> 
                    {colorOptions && colorOptions.map(
                      opt => <option key={opt.ID} value={`${opt.ID}`}>{opt.Name}</option> 
                    )}
                  </Select>
                </FormControl>
                <FormControl>
                  <FormLabel>Build Date</FormLabel>
                  <Input name='buildDate' type='datetime-local' defaultValue={`${getDefaultDate()}`} />      
                </FormControl>
              </VStack>
            </form>
          </ModalBody>
          <ModalFooter>
            <Button colorScheme='blue' variant={'outline'} mr={3} onClick={onClose}> Cancel </Button>
            <Button colorScheme='blue' onClick={() => formRef.current.dispatchEvent(
                new Event("submit", { cancelable: true, bubbles: true })
            )}> Save </Button>            
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  )
}

function DeleteConfirmModal({car}) {
  const { isOpen, onOpen, onClose } = useDisclosure()

  return (
    <>
      <Button colorScheme='red' onClick={onOpen} variant={'outline'}>
        Delete car
      </Button>

      <AlertDialog
        isOpen={isOpen}        
        onClose={onClose}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize='lg' fontWeight='bold'>
              Delete Car: {car?.make}
            </AlertDialogHeader>

            <AlertDialogBody>
              Are you sure? You can't undo this action afterwards.
            </AlertDialogBody>

            <AlertDialogFooter>
              <Button onClick={onClose}>
                Cancel
              </Button>
              <Button colorScheme='red' onClick={onClose} ml={3}>
                Delete
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
    </>
  )
}

function App() {
  const [cars, setCars] = useState([]);
  const [colors, setColors] = useState([]);

  async function getCars() {
    const response = await fetch("/api/car/cars");
    const cars = await response.json();

    if (cars?.result.length > 0) {
      return cars.result;
    }

    return [];
  }

  async function getColors() {
    const response = await fetch("/api/color/colors");
    const colors = await response.json();

    if (colors?.result.length > 0) {
      return colors.result;
    }

    return [];
  }

  useEffect(() => {
    getCars().then(r => {      
      setCars(r);
    });

    getColors().then(r => {      
      setColors(r);
    });
  }, []);
  

  return (
    <ChakraProvider theme={theme}>
      <Box textAlign="center" fontSize="xl">
        <Flex minH="100vh" p={5} flexDir={"column"} gap={3}>
          <Flex justifyContent={"flex-end"}>
            <ColorModeSwitcher />
          </Flex>          
          <Card>
            <CardHeader display={"flex"} justifyContent={"space-between"}>
              Available cars
              <CarFormModal colorOptions={colors} setCars={setCars} />
            </CardHeader>
            <CardBody>
              <VStack spacing={4} divider={<StackDivider />}>   
                {cars.length && cars.map((car, index) => {
                  return (
                    <Flex key={`car-${index}`} justifyContent={"flex-start"} flexDir={"column"} w={"full"}> 
                      <Flex gap={3}>
                        <Text as={"b"}>
                          Make:
                        </Text>
                        {car.make}                        
                      </Flex>  
                      <Flex gap={3}>
                        <Text as={"b"}>
                          Model:
                        </Text>
                        {car.model}
                      </Flex>  
                      <Flex gap={3}>
                        <Text as={"b"}>
                          Color:
                        </Text>
                        {car.color}
                      </Flex>  
                      <Flex gap={3}>
                        <Text as={"b"}>
                          Build Date:
                        </Text>
                        {new Date(car.buildDate).toLocaleDateString()}
                      </Flex>                       
                      <Flex mt={3} gap={3}>                        
                        <CarFormModal isEdit car={car} colorOptions={colors}/>
                        <DeleteConfirmModal car={car} />
                      </Flex>
                    </Flex>
                  )
                })}   
                {cars.length <= 0 && "You have no available cars at the moment"}             
              </VStack>
            </CardBody>
          </Card>          
        </Flex>
      </Box>
    </ChakraProvider>
  );
}

export default App;
