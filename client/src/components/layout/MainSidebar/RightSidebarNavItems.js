import React from "react";
import {Nav, Button} from "shards-react";
import ChildSidebarNavItem from "./ChildSidebarNavItem";
import TreeChildbarNavItem from "./TreeChildbarNavItem";
import Store from "../../../flux/store";

class RightSidebarNavItems extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            navItems: Store.getSidebarItems(),
            resourceDetails: Store.getDetailedResourceTableData()
        };

        this.onChange = this.onChange.bind(this);
    }

    componentWillMount() {
        Store.addChangeListener(this.onChange);
    }

    componentWillUnmount() {
        Store.removeChangeListener(this.onChange);
    }

    onChange() {
        this.setState({
            ...this.state,
            navItems: Store.getSidebarItems(),
            resourceDetails: Store.getDetailedResourceTableData()
        });
    }

    render() {
        const {navItems: items, resourceDetails: resData} = this.state;
        const selectedItem = items.filter(item => item.to === this.props.location.pathname).shift();
        const isManageResources = (selectedItem.to === "/manage-resources");
        return (
            <div className="child-nav-wrapper">
                <h6 className="ml-10 mt-10 font-weight-bold text-black">{selectedItem.name}</h6>
                <hr className="mx-10"/>
                <Button squared className="w-100 mb-4">
                    <i className="fa fa-plus mr-1"></i>
                    Add Resources
                </Button>
                <Nav className="nav--no-borders flex-column">
                    {isManageResources ? <TreeChildbarNavItem data={resData}/> :
                        selectedItem.children && selectedItem.children.map((item, idx) => {
                            return <ChildSidebarNavItem key={idx} item={item}/>
                        })}
                </Nav>
            </div>
        )
    }
}

export default RightSidebarNavItems;
