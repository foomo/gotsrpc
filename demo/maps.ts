

interface CarMap {
    [index:string]:Car;
}

interface Car {
    color:string;
    wheels:number;
    extras?:{[index:string]:string}
}

interface Complex {
    cars:CarMap;
    people:{[index:string]:{
        [index:string]:Car
    }};
}

var foo : Complex = {
    people: {
        bla:{
            audi: {
                color: "white",
                wheels:1
            }
        }
    },
    cars: {
        audi: {
            color: "white",
            wheels:1,
            extras:{
                lights:"xenon"
            }
        },
        blas: {
            color:"red",
            wheels:23.1
        }
    }
}